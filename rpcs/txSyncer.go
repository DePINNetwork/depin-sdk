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
	"fmt"
	"sync"
	"time"

	"github.com/DePINNetwork/depin-sdk/data"
	"github.com/DePINNetwork/depin-sdk/data/transactions"
	"github.com/DePINNetwork/depin-sdk/logging"
	"github.com/DePINNetwork/depin-sdk/network"
	"github.com/DePINNetwork/depin-sdk/util/bloom"
)

// PendingTxAggregate is a container of pending transactions
type PendingTxAggregate interface {
	PendingTxIDs() []transactions.Txid
	PendingTxGroups() [][]transactions.SignedTxn
}

// TxSyncClient abstracts sync-ing pending transactions from a peer.
type TxSyncClient interface {
	Sync(ctx context.Context, bloom *bloom.Filter) (txns [][]transactions.SignedTxn, err error)
	Address() string
	Close() error
}

// TxSyncer fetches pending transactions that are missing from its pool, and feeds them to the handler
type TxSyncer struct {
	pool         PendingTxAggregate
	clientSource network.GossipNode
	handler      data.SolicitedTxHandler
	ctx          context.Context
	cancel       context.CancelFunc
	syncInterval time.Duration
	syncTimeout  time.Duration
	counter      uint32
	wg           sync.WaitGroup
	log          logging.Logger
	httpSync     *HTTPTxSync
}

// MakeTxSyncer returns a TxSyncer
func MakeTxSyncer(pool PendingTxAggregate, clientSource network.GossipNode, txHandler data.SolicitedTxHandler, syncInterval time.Duration, syncTimeout time.Duration, serverResponseSize int) *TxSyncer {
	return &TxSyncer{
		pool:         pool,
		clientSource: clientSource,
		handler:      txHandler,
		syncInterval: syncInterval,
		syncTimeout:  syncTimeout,
		log:          logging.Base(),
		httpSync:     makeHTTPSync(clientSource, logging.Base(), uint64(serverResponseSize)),
	}
}

// Start begins periodically syncing after the canStart chanel indicates it can begin
func (syncer *TxSyncer) Start(canStart chan struct{}) {
	syncer.wg.Add(1)
	syncer.ctx, syncer.cancel = context.WithCancel(context.Background())
	go func() {
		defer syncer.wg.Done()
		select {
		case <-syncer.ctx.Done():
			return
		case <-canStart:
		}
		for {
			select {
			case <-syncer.ctx.Done():
				return
			case <-time.After(syncer.syncInterval):
				err := syncer.sync()
				if err != nil {
					syncer.log.Warnf("problem syncing transactions %v", err)
				}
			}
		}
	}()
}

// Stop stops periodic syncing
func (syncer *TxSyncer) Stop() {
	syncer.log.Debug("transaction syncer is stopping")
	defer syncer.log.Debug("transaction syncer has stopped")

	syncer.cancel()
	syncer.wg.Wait()
}

func (syncer *TxSyncer) sync() error {
	return syncer.syncFromClient(syncer.httpSync)
}

const bloomFilterFalsePositiveRate = 0.01

func (syncer *TxSyncer) syncFromClient(client TxSyncClient) error {
	syncer.log.Infof("TxSyncer.Sync: asking client %v for missing transactions", client.Address())

	pending := syncer.pool.PendingTxIDs()
	sizeBits, numHashes := bloom.Optimal(len(pending), bloomFilterFalsePositiveRate)
	filter := bloom.New(sizeBits, numHashes, syncer.counter)
	syncer.counter++
	for _, txid := range pending {
		filter.Set(txid[:])
	}

	ctx, cf := context.WithTimeout(syncer.ctx, syncer.syncTimeout)
	defer cf()
	txgroups, err := client.Sync(ctx, filter)
	if err != nil {
		return fmt.Errorf("TxSyncer.Sync: peer '%v' error '%v'", client.Address(), err)
	}

	var pendingTxidMap map[transactions.Txid]struct{}
	// test to see if all the transaction that we've received honor the bloom filter constraints
	// that we've requested.
	for _, txgroup := range txgroups {
		var txnsInFilter int
		for i := range txgroup {
			txID := txgroup[i].ID()
			if filter.Test(txID[:]) {
				// having the transaction id tested here might still fall into the false-positive class, so we
				// need to perform explicit check. This is not too bad since we're doing this check only on the fail
				// cases.
				if pendingTxidMap == nil {
					// construct and initialize it.
					pendingTxidMap = make(map[transactions.Txid]struct{}, len(pending))
					for _, txid := range pending {
						pendingTxidMap[txid] = struct{}{}
					}
				}
				if _, has := pendingTxidMap[txID]; has {
					// we just found a transaction that shouldn't have been
					// included in the response.  maybe this is a false positive
					// and other transactions in the group aren't included in the
					// bloom filter, though.
					txnsInFilter++
				}
			}
		}

		// if the entire group was in the bloom filter, report an error.
		if txnsInFilter == len(txgroup) {
			client.Close()
			return fmt.Errorf("TxSyncer.Sync: peer %v sent a transaction group that was entirely included in the bloom filter", client.Address())
		}

		// send the transaction to the trasaction pool
		if syncer.handler.Handle(txgroup) != nil {
			client.Close()
			return fmt.Errorf("TxSyncer.Sync: peer %v sent invalid transaction", client.Address())
		}
	}

	return nil
}
