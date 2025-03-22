package store

import (
	corestore "cosmossdk.io/core/store"
	"github.com/depinnetwork/depin-sdk/store/cache"
	"github.com/depinnetwork/depin-sdk/store/metrics"
	"github.com/depinnetwork/depin-sdk/store/rootmulti"
	"github.com/depinnetwork/depin-sdk/store/types"
)

func NewCommitMultiStore(db corestore.KVStoreWithBatch, logger types.Logger, metricGatherer metrics.StoreMetrics) types.CommitMultiStore {
	return rootmulti.NewStore(db, logger, metricGatherer)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
