package rootmulti

import (
	"github.com/depinnetwork/depin-sdk/store/dbadapter"
	pruningtypes "github.com/depinnetwork/depin-sdk/store/pruning/types"
	"github.com/depinnetwork/depin-sdk/store/types"
)

var commithash = []byte("FAKE_HASH")

var (
	_ types.KVStore   = (*commitDBStoreAdapter)(nil)
	_ types.Committer = (*commitDBStoreAdapter)(nil)
)

//----------------------------------------
// commitDBStoreWrapper should only be used for simulation/debugging,
// as it doesn't compute any commit hash, and it cannot load older state.

// Wrapper type for corestore.KVStoreWithBatch with implementation of KVStore
type commitDBStoreAdapter struct {
	dbadapter.Store
}

func (cdsa commitDBStoreAdapter) Commit() types.CommitID {
	return types.CommitID{
		Version: -1,
		Hash:    commithash,
	}
}

func (cdsa commitDBStoreAdapter) LastCommitID() types.CommitID {
	return types.CommitID{
		Version: -1,
		Hash:    commithash,
	}
}

func (cdsa commitDBStoreAdapter) LatestVersion() int64 {
	return -1
}

func (cdsa commitDBStoreAdapter) WorkingHash() []byte {
	return commithash
}

func (cdsa commitDBStoreAdapter) SetPruning(_ pruningtypes.PruningOptions) {}

// GetPruning is a no-op as pruning options cannot be directly set on this store.
// They must be set on the root commit multi-store.
func (cdsa commitDBStoreAdapter) GetPruning() pruningtypes.PruningOptions {
	return pruningtypes.NewPruningOptions(pruningtypes.PruningUndefined)
}
