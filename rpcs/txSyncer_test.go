// Copyright (C) 2019-2025 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package rpcs

import (
	"context"
	"errors"
	"math/rand"
	"net/http"
	"net/rpc"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/DePINNetwork/depin-sdk/components/mocks"
	"github.com/DePINNetwork/depin-sdk/config"
	"github.com/DePINNetwork/depin-sdk/crypto"
	"github.com/DePINNetwork/depin-sdk/data/bookkeeping"
	"github.com/DePINNetwork/depin-sdk/data/transactions"
	"github.com/DePINNetwork/depin-sdk/logging"
	"github.com/DePINNetwork/depin-sdk/network"
	"github.com/DePINNetwork/depin-sdk/protocol"
	"github.com/DePINNetwork/depin-sdk/test/partitiontest"
	"github.com/DePINNetwork/depin-sdk/util/bloom"
)

type mockPendingTxAggregate struct {
	txns []transactions.SignedTxn
}

var testSource rand.Source
var testRand *rand.Rand

func init() {
	testSource = rand.NewSource(12345678)
	testRand = rand.New(testSource)
}

func testRandBytes(d []byte) {
	// We don't need cryptographically strong random bytes for a
	// unit test, we _do_ need deterministic 'random' bytes so
	// that _sometimes_ a bloom filter doesn't fail on the data
	// (e.g. TestSync() below).
	n, err := testRand.Read(d)
	if n != len(d) {
		panic("short rand read")
	}
	if err != nil {
		panic(err)
	}
}

func makeMockPendingTxAggregate(txCount int) mockPendingTxAggregate {
	var secret [32]byte
	testRandBytes(secret[:])
	sk := crypto.GenerateSignatureSecrets(crypto.Seed(secret))
	mock := mockPendingTxAggregate{
		txns: make([]transactions.SignedTxn, txCount),
	}

	for i := 0; i < txCount; i++ {
		var note [16]byte
		testRandBytes(note[:])
		tx := transactions.Transaction{
			Type: protocol.PaymentTx,
			Header: transactions.Header{
				Note: note[:],
			},
		}
		stx := tx.Sign(sk)
		mock.txns[i] = stx
	}
	return mock
}

func (mock mockPendingTxAggregate) PendingTxIDs() []transactions.Txid {
	// return all but one ID
	ids := make([]transactions.Txid, 0)
	for _, tx := range mock.txns {
		ids = append(ids, tx.ID())
	}
	return ids
}
func (mock mockPendingTxAggregate) PendingTxGroups() [][]transactions.SignedTxn {
	return bookkeeping.SignedTxnsToGroups(mock.txns)
}

type mockHandler struct {
	messageCounter atomic.Int32
	err            error
}

func (handler *mockHandler) Handle(txgroup []transactions.SignedTxn) error {
	handler.messageCounter.Add(1)
	return handler.err
}

const testSyncInterval = 5 * time.Second
const testSyncTimeout = 4 * time.Second

type mockRunner struct {
	ran           bool
	done          chan *rpc.Call
	failWithNil   bool
	failWithError bool
	txgroups      [][]transactions.SignedTxn
}

type mockRPCClient struct {
	client  *mockRunner
	closed  bool
	rootURL string
	log     logging.Logger
}

func (client *mockRPCClient) Close() error {
	client.closed = true
	return nil
}

func (client *mockRPCClient) Address() string {
	return "mock.address."
}
func (client *mockRPCClient) Sync(ctx context.Context, bloom *bloom.Filter) (txgroups [][]transactions.SignedTxn, err error) {
	client.log.Info("MockRPCClient.Sync")
	select {
	case <-ctx.Done():
		return nil, errors.New("cancelled")
	default:
	}
	if client.client.failWithNil {
		return nil, errors.New("old failWithNil")
	}
	if client.client.failWithError {
		return nil, errors.New("failing call")
	}
	return client.client.txgroups, nil
}

// network.HTTPPeer interface
func (client *mockRPCClient) GetAddress() string {
	return client.rootURL
}

func (client *mockRPCClient) GetHTTPClient() *http.Client {
	return nil
}

type mockClientAggregator struct {
	mocks.MockNetwork
	peers []network.Peer
}

func (mca *mockClientAggregator) GetPeers(options ...network.PeerOption) []network.Peer {
	return mca.peers
}

func (mca *mockClientAggregator) GetHTTPClient(address string) (*http.Client, error) {
	return &http.Client{
		Transport: &network.HTTPPAddressBoundTransport{
			Addr:           address,
			InnerTransport: http.DefaultTransport},
	}, nil
}

func TestSyncFromClient(t *testing.T) {
	partitiontest.PartitionTest(t)

	clientPool := makeMockPendingTxAggregate(2)
	serverPool := makeMockPendingTxAggregate(1)
	runner := mockRunner{failWithNil: false, failWithError: false, txgroups: serverPool.PendingTxGroups()[len(serverPool.PendingTxGroups())-1:], done: make(chan *rpc.Call)}
	client := mockRPCClient{client: &runner, log: logging.TestingLog(t)}
	clientAgg := mockClientAggregator{peers: []network.Peer{&client}}
	handler := mockHandler{}
	syncer := MakeTxSyncer(clientPool, &clientAgg, &handler, testSyncInterval, testSyncTimeout, config.GetDefaultLocal().TxSyncServeResponseSize)
	// Since syncer is not Started, set the context here
	syncer.ctx, syncer.cancel = context.WithCancel(context.Background())
	syncer.log = logging.TestingLog(t)

	require.NoError(t, syncer.syncFromClient(&client))
	require.Equal(t, int32(1), handler.messageCounter.Load())
}

func TestSyncFromUnsupportedClient(t *testing.T) {
	partitiontest.PartitionTest(t)

	pool := makeMockPendingTxAggregate(3)
	runner := mockRunner{failWithNil: true, failWithError: false, txgroups: pool.PendingTxGroups()[len(pool.PendingTxGroups())-1:], done: make(chan *rpc.Call)}
	client := mockRPCClient{client: &runner, log: logging.TestingLog(t)}
	clientAgg := mockClientAggregator{peers: []network.Peer{&client}}
	handler := mockHandler{}
	syncer := MakeTxSyncer(pool, &clientAgg, &handler, testSyncInterval, testSyncTimeout, config.GetDefaultLocal().TxSyncServeResponseSize)
	// Since syncer is not Started, set the context here
	syncer.ctx, syncer.cancel = context.WithCancel(context.Background())
	syncer.log = logging.TestingLog(t)

	require.Error(t, syncer.syncFromClient(&client))
	require.Zero(t, handler.messageCounter.Load())
}

func TestSyncFromClientAndQuit(t *testing.T) {
	partitiontest.PartitionTest(t)

	pool := makeMockPendingTxAggregate(3)
	runner := mockRunner{failWithNil: false, failWithError: false, txgroups: pool.PendingTxGroups()[len(pool.PendingTxGroups())-1:], done: make(chan *rpc.Call)}
	client := mockRPCClient{client: &runner, log: logging.TestingLog(t)}
	clientAgg := mockClientAggregator{peers: []network.Peer{&client}}
	handler := mockHandler{}
	syncer := MakeTxSyncer(pool, &clientAgg, &handler, testSyncInterval, testSyncTimeout, config.GetDefaultLocal().TxSyncServeResponseSize)
	// Since syncer is not Started, set the context here
	syncer.ctx, syncer.cancel = context.WithCancel(context.Background())
	syncer.log = logging.TestingLog(t)
	syncer.cancel()
	require.Error(t, syncer.syncFromClient(&client))
	require.Zero(t, handler.messageCounter.Load())
}

func TestSyncFromClientAndError(t *testing.T) {
	partitiontest.PartitionTest(t)

	pool := makeMockPendingTxAggregate(3)
	runner := mockRunner{failWithNil: false, failWithError: true, txgroups: pool.PendingTxGroups()[len(pool.PendingTxGroups())-1:], done: make(chan *rpc.Call)}
	client := mockRPCClient{client: &runner, log: logging.TestingLog(t)}
	clientAgg := mockClientAggregator{peers: []network.Peer{&client}}
	handler := mockHandler{}
	syncer := MakeTxSyncer(pool, &clientAgg, &handler, testSyncInterval, testSyncTimeout, config.GetDefaultLocal().TxSyncServeResponseSize)
	// Since syncer is not Started, set the context here
	syncer.ctx, syncer.cancel = context.WithCancel(context.Background())
	syncer.log = logging.TestingLog(t)
	require.Error(t, syncer.syncFromClient(&client))
	require.Zero(t, handler.messageCounter.Load())
}

func TestSyncFromClientAndTimeout(t *testing.T) {
	partitiontest.PartitionTest(t)

	pool := makeMockPendingTxAggregate(3)
	runner := mockRunner{failWithNil: false, failWithError: false, txgroups: pool.PendingTxGroups()[len(pool.PendingTxGroups())-1:], done: make(chan *rpc.Call)}
	client := mockRPCClient{client: &runner, log: logging.TestingLog(t)}
	clientAgg := mockClientAggregator{peers: []network.Peer{&client}}
	handler := mockHandler{}
	syncTimeout := time.Duration(0)
	syncer := MakeTxSyncer(pool, &clientAgg, &handler, testSyncInterval, syncTimeout, config.GetDefaultLocal().TxSyncServeResponseSize)
	// Since syncer is not Started, set the context here
	syncer.ctx, syncer.cancel = context.WithCancel(context.Background())
	syncer.log = logging.TestingLog(t)
	require.Error(t, syncer.syncFromClient(&client))
	require.Zero(t, handler.messageCounter.Load())
}

func TestSync(t *testing.T) {
	partitiontest.PartitionTest(t)

	pool := makeMockPendingTxAggregate(1)
	nodeA := basicRPCNode{}
	txservice := makeTxService(pool, "test genesisID", config.GetDefaultLocal().TxPoolSize, config.GetDefaultLocal().TxSyncServeResponseSize)
	nodeA.RegisterHTTPHandler(TxServiceHTTPPath, txservice)
	nodeA.start()
	nodeAURL := nodeA.rootURL()

	runner := mockRunner{failWithNil: false, failWithError: false, txgroups: pool.PendingTxGroups()[len(pool.PendingTxGroups())-1:], done: make(chan *rpc.Call)}
	client := mockRPCClient{client: &runner, rootURL: nodeAURL, log: logging.TestingLog(t)}
	clientAgg := mockClientAggregator{peers: []network.Peer{&client}, MockNetwork: mocks.MockNetwork{GenesisID: "test genesisID"}}
	handler := mockHandler{}
	syncerPool := makeMockPendingTxAggregate(3)
	syncer := MakeTxSyncer(syncerPool, &clientAgg, &handler, testSyncInterval, testSyncTimeout, config.GetDefaultLocal().TxSyncServeResponseSize)
	// Since syncer is not Started, set the context here
	syncer.ctx, syncer.cancel = context.WithCancel(context.Background())
	syncer.log = logging.TestingLog(t)

	require.NoError(t, syncer.sync())
	require.Equal(t, int32(1), handler.messageCounter.Load())
}

func TestNoClientsSync(t *testing.T) {
	partitiontest.PartitionTest(t)

	pool := makeMockPendingTxAggregate(3)
	clientAgg := mockClientAggregator{peers: []network.Peer{}}
	handler := mockHandler{}
	syncer := MakeTxSyncer(pool, &clientAgg, &handler, testSyncInterval, testSyncTimeout, config.GetDefaultLocal().TxSyncServeResponseSize)
	// Since syncer is not Started, set the context here
	syncer.ctx, syncer.cancel = context.WithCancel(context.Background())
	syncer.log = logging.TestingLog(t)

	require.NoError(t, syncer.sync())
	require.Zero(t, handler.messageCounter.Load())
}

func TestStartAndStop(t *testing.T) {
	partitiontest.PartitionTest(t)

	pool := makeMockPendingTxAggregate(1)
	nodeA := basicRPCNode{}
	txservice := makeTxService(pool, "test genesisID", config.GetDefaultLocal().TxPoolSize, config.GetDefaultLocal().TxSyncServeResponseSize)
	nodeA.RegisterHTTPHandler(TxServiceHTTPPath, txservice)
	nodeA.start()
	nodeAURL := nodeA.rootURL()

	runner := mockRunner{failWithNil: false, failWithError: false, txgroups: pool.PendingTxGroups()[len(pool.PendingTxGroups())-1:], done: make(chan *rpc.Call)}
	client := mockRPCClient{client: &runner, rootURL: nodeAURL, log: logging.TestingLog(t)}
	clientAgg := mockClientAggregator{peers: []network.Peer{&client}, MockNetwork: mocks.MockNetwork{GenesisID: "test genesisID"}}
	handler := mockHandler{}

	syncerPool := makeMockPendingTxAggregate(0)
	syncInterval := time.Second
	syncTimeout := time.Second
	syncer := MakeTxSyncer(syncerPool, &clientAgg, &handler, syncInterval, syncTimeout, config.GetDefaultLocal().TxSyncServeResponseSize)
	syncer.log = logging.TestingLog(t)

	// ensure that syncing doesn't start
	canStart := make(chan struct{})
	syncer.Start(canStart)
	time.Sleep(2 * time.Second)
	require.Zero(t, handler.messageCounter.Load())

	// signal that syncing can start
	close(canStart)
	for x := 0; x < 20; x++ {
		time.Sleep(100 * time.Millisecond)
		if handler.messageCounter.Load() != 0 {
			break
		}
	}
	require.Equal(t, int32(1), handler.messageCounter.Load())

	// stop syncing and ensure it doesn't happen
	syncer.Stop()
	time.Sleep(2 * time.Second)
	require.Equal(t, int32(1), handler.messageCounter.Load())
}

func TestStartAndQuit(t *testing.T) {
	partitiontest.PartitionTest(t)

	pool := makeMockPendingTxAggregate(3)
	runner := mockRunner{failWithNil: false, failWithError: false, txgroups: pool.PendingTxGroups()[len(pool.PendingTxGroups())-1:], done: make(chan *rpc.Call)}
	client := mockRPCClient{client: &runner, log: logging.TestingLog(t)}
	clientAgg := mockClientAggregator{peers: []network.Peer{&client}}
	handler := mockHandler{}
	syncInterval := time.Second
	syncTimeout := time.Second
	syncer := MakeTxSyncer(pool, &clientAgg, &handler, syncInterval, syncTimeout, config.GetDefaultLocal().TxSyncServeResponseSize)
	syncer.log = logging.TestingLog(t)

	// ensure that syncing doesn't start
	canStart := make(chan struct{})
	syncer.Start(canStart)
	time.Sleep(2 * time.Second)
	require.Zero(t, handler.messageCounter.Load())

	syncer.cancel()
	time.Sleep(50 * time.Millisecond)
	// signal that syncing can start, but ensure that it doesn't start (since we quit)
	close(canStart)
	time.Sleep(2 * time.Second)
	require.Zero(t, handler.messageCounter.Load())
}
