package store

import (
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/v43/store/cache"
	"github.com/cosmos/cosmos-sdk/v43/store/rootmulti"
	"github.com/cosmos/cosmos-sdk/v43/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
