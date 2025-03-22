package transient

import (
	coretesting "cosmossdk.io/core/testing"
	"github.com/depinnetwork/depin-sdk/store/dbadapter"
	pruningtypes "github.com/depinnetwork/depin-sdk/store/pruning/types"
	"github.com/depinnetwork/depin-sdk/store/types"
)

var (
	_ types.Committer = (*Store)(nil)
	_ types.KVStore   = (*Store)(nil)
)

// Store is a wrapper for a MemDB with Committer implementation
type Store struct {
	dbadapter.Store
}

// NewStore constructs new MemDB adapter
func NewStore() *Store {
	return &Store{Store: dbadapter.Store{DB: coretesting.NewMemDB()}}
}

// Commit cleans up Store.
// Implements CommitStore
func (ts *Store) Commit() (id types.CommitID) {
	ts.Store = dbadapter.Store{DB: coretesting.NewMemDB()}
	return
}

func (ts *Store) SetPruning(_ pruningtypes.PruningOptions) {}

// GetPruning is a no-op as pruning options cannot be directly set on this store.
// They must be set on the root commit multi-store.
func (ts *Store) GetPruning() pruningtypes.PruningOptions {
	return pruningtypes.NewPruningOptions(pruningtypes.PruningUndefined)
}

// LastCommitID implements CommitStore
func (ts *Store) LastCommitID() types.CommitID {
	return types.CommitID{}
}

// LatestVersion implements Committer
func (ts *Store) LatestVersion() int64 {
	return 0
}

func (ts *Store) WorkingHash() []byte {
	return []byte{}
}

// GetStoreType implements Store.
func (ts *Store) GetStoreType() types.StoreType {
	return types.StoreTypeTransient
}
