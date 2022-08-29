package store

import (
	"github.com/commit-smart-core-banking-system/store/postgres_store"
)

type DataStoreOptions struct {
	DbDriver string
	DbSource string
}

var Store DataStore

var availDataStores = map[string]func(options DataStoreOptions) NewDataStoreClientResp{
	"postgres": func(options DataStoreOptions) NewDataStoreClientResp {
		client, err := postgres_store.NewClient(options.DbSource)
		if err != nil {
			return NewDataStoreClientResp{Store: client, Error: err}
		}
		return NewDataStoreClientResp{Store: client, Error: nil}
	},
}

type NewDataStoreClientResp struct {
	Store DataStore
	Error error
}

//Initialize new store based on db driver
func NewDataStoreClient(options DataStoreOptions) NewDataStoreClientResp {
	return availDataStores[options.DbDriver](options)
}
