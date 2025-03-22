package adapters

import (
	"context"
	
	cmtabciv1 "github.com/cometbft/cometbft/api/cometbft/abci/v1"
	depinabciv1 "github.com/depinnetwork/por-consensus/api/cometbft/abci/v1"
	cmtproto "github.com/depinnetwork/por-consensus/api/cometbft/abci/v1"
	storetypes "cosmossdk.io/store/types"
)

// CometBFTABCIListenerWrapper wraps a DePIN ABCIListener to be used with CometBFT interfaces
type CometBFTABCIListenerWrapper struct {
	Listener interface {
		ListenCommit(context.Context, depinabciv1.CommitResponse, []*storetypes.StoreKVPair) error
		ListenFinalizeBlock(context.Context, depinabciv1.FinalizeBlockRequest, depinabciv1.FinalizeBlockResponse) error
	}
}

// ListenCommit adapts between different CommitResponse types
func (a *CometBFTABCIListenerWrapper) ListenCommit(
	ctx context.Context,
	resp cmtabciv1.CommitResponse,
	kvs []*storetypes.StoreKVPair,
) error {
	// Convert cometbft CommitResponse to depinnetwork CommitResponse
	depinResp := CometBFTToDepinCommitResponse(resp)
	
	// Call the inner listener
	return a.Listener.ListenCommit(ctx, depinResp, kvs)
}

// ListenFinalizeBlock adapts between different FinalizeBlock types
func (a *CometBFTABCIListenerWrapper) ListenFinalizeBlock(
	ctx context.Context,
	req cmtabciv1.FinalizeBlockRequest,
	res cmtabciv1.FinalizeBlockResponse,
) error {
	// Convert to the depin expected types
	depinReq := CometBFTToDepinFinalizeBlockRequest(req)
	depinRes := CometBFTToDepinFinalizeBlockResponse(res)
	
	// Call the inner listener
	return a.Listener.ListenFinalizeBlock(ctx, depinReq, depinRes)
}

// DepinABCIListenerWrapper wraps a CometBFT ABCIListener to be used with DePIN interfaces
type DepinABCIListenerWrapper struct {
	Listener interface {
		ListenCommit(context.Context, cmtabciv1.CommitResponse, []*storetypes.StoreKVPair) error
		ListenFinalizeBlock(context.Context, cmtabciv1.FinalizeBlockRequest, cmtabciv1.FinalizeBlockResponse) error
	}
}

// ListenCommit adapts between different CommitResponse types
func (a *DepinABCIListenerWrapper) ListenCommit(
	ctx context.Context,
	resp depinabciv1.CommitResponse,
	kvs []*storetypes.StoreKVPair,
) error {
	// Convert depin CommitResponse to cometbft CommitResponse
	cmtResp := DepinToCometBFTCommitResponse(resp)
	
	// Call the inner listener
	return a.Listener.ListenCommit(ctx, cmtResp, kvs)
}

// ListenFinalizeBlock adapts between different FinalizeBlock types
func (a *DepinABCIListenerWrapper) ListenFinalizeBlock(
	ctx context.Context,
	req depinabciv1.FinalizeBlockRequest,
	res depinabciv1.FinalizeBlockResponse,
) error {
	// Convert to the cometbft expected types
	cmtReq := DepinToCometBFTFinalizeBlockRequest(req)
	cmtRes := DepinToCometBFTFinalizeBlockResponse(res)
	
	// Call the inner listener
	return a.Listener.ListenFinalizeBlock(ctx, cmtReq, cmtRes)
}

// NewCometBFTABCIListenerWrapper creates a new adapter for DePIN ABCIListener
func NewCometBFTABCIListenerWrapper(listener interface{}) *CometBFTABCIListenerWrapper {
	return &CometBFTABCIListenerWrapper{Listener: listener}
}

// NewDepinABCIListenerWrapper creates a new adapter for CometBFT ABCIListener
func NewDepinABCIListenerWrapper(listener interface{}) *DepinABCIListenerWrapper {
	return &DepinABCIListenerWrapper{Listener: listener}
}

// CreateABCIListenerAdapter determines whether the passed in listener implements
// DePIN or CometBFT ABCIListener, and returns the appropriate wrapper
func CreateABCIListenerAdapter(listener interface{}) interface{} {
	// Check if it's a DePIN listener
	if _, ok := listener.(interface {
		ListenCommit(context.Context, depinabciv1.CommitResponse, []*storetypes.StoreKVPair) error
	}); ok {
		return NewCometBFTABCIListenerWrapper(listener)
	}
	
	// Check if it's a CometBFT listener
	if _, ok := listener.(interface {
		ListenCommit(context.Context, cmtabciv1.CommitResponse, []*storetypes.StoreKVPair) error
	}); ok {
		return NewDepinABCIListenerWrapper(listener)
	}
	
	// Return original if it doesn't match known patterns
	return listener
}
