package main

import (
	"github.com/commit-smart-core-banking-system/config"
	"github.com/commit-smart-core-banking-system/logger"
	"github.com/commit-smart-core-banking-system/server"
	"github.com/commit-smart-core-banking-system/store"
)

func init() {
	//Define Log Level
	logger.NewLogger(logger.Config{
		Service: "domain",
		Level:   "debug",
	})

	err := config.LoadConfig(".")
	if err != nil {
		logger.Panic("Unable to load app config")
	}

	storeResp := store.NewDataStoreClient(store.DataStoreOptions{DbDriver: config.AppConfiguration.DbDriver, DbSource: config.AppConfiguration.DbSource})
	if storeResp.Error != nil {
		logger.Panic("Unable to connect to store", logger.LogErrorField(storeResp.Error))
	}
	store.Store = storeResp.Store
}

func main() {

	if err := server.NewServer().Run(); err != nil {
		logger.Panic("Unable to run server", logger.LogErrorField(err))
	}

}
