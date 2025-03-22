package mock

import "github.com/depinnetwork/depin-sdk/store/v2"

// StateCommitter is a mock of store.Committer
type StateCommitter interface {
	store.Committer
	store.Pruner
	store.PausablePruner
	store.UpgradeableStore
	store.VersionedReader
	store.UpgradableDatabase
}
