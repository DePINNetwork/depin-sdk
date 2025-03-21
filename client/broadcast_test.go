package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/depinnetwork/por-consensus/crypto/tmhash"
	"github.com/depinnetwork/por-consensus/mempool"
	"github.com/depinnetwork/por-consensus/rpc/client/mock"
	coretypes "github.com/depinnetwork/por-consensus/rpc/core/types"
	cmttypes "github.com/depinnetwork/por-consensus/types"
	"github.com/stretchr/testify/require"

	"github.com/depinnetwork/depin-sdk/client/flags"
	sdkerrors "github.com/depinnetwork/depin-sdk/types/errors"
)

type MockClient struct {
	mock.Client
	err error
}

func (c MockClient) BroadcastTxAsync(_ context.Context, _ cmttypes.Tx) (*coretypes.ResultBroadcastTx, error) {
	return nil, c.err
}

func (c MockClient) BroadcastTxSync(_ context.Context, _ cmttypes.Tx) (*coretypes.ResultBroadcastTx, error) {
	return nil, c.err
}

func CreateContextWithErrorAndMode(err error, mode string) Context {
	return Context{
		Client:        MockClient{err: err},
		BroadcastMode: mode,
	}
}

// Test the correct code is returned when
func TestBroadcastError(t *testing.T) {
	errors := map[error]uint32{
		mempool.ErrTxInCache:       sdkerrors.ErrTxInMempoolCache.ABCICode(),
		mempool.ErrTxTooLarge{}:    sdkerrors.ErrTxTooLarge.ABCICode(),
		mempool.ErrMempoolIsFull{}: sdkerrors.ErrMempoolIsFull.ABCICode(),
	}

	modes := []string{
		flags.BroadcastAsync,
		flags.BroadcastSync,
	}

	txBytes := []byte{0xA, 0xB}
	txHash := fmt.Sprintf("%X", tmhash.Sum(txBytes))

	for _, mode := range modes {
		for err, code := range errors {
			ctx := CreateContextWithErrorAndMode(err, mode)
			resp, returnedErr := ctx.BroadcastTx(txBytes)
			require.NoError(t, returnedErr)
			require.Equal(t, code, resp.Code)
			require.NotEmpty(t, resp.Codespace)
			require.Equal(t, txHash, resp.TxHash)
		}
	}
}
