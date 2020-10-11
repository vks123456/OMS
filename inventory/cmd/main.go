package main

import (
	cfg "OMS/inventory/config"
	"OMS/inventory/internal/cache"
	"OMS/inventory/internal/db"
	"OMS/inventory/internal/httpserver"

	"github.com/allegro/bigcache"
	"github.com/rs/zerolog/log"
)

func main() {
	//Initialise Configs
	config := cfg.InitConfigStore()
	log.Info().Msg("Starting Inventory service...")

	//Initialise mysql db
	log.Info().Msg("Initializing Mysql db...")
	mysql := &db.MysqlDB{}
	mysql.InitializeMySQL()

	//Initialise local cache
	log.Info().Msg("Initializing App Cache...")
	cacheProvider := &cache.AppCache{}
	if err := cacheProvider.Initialize(bigcache.DefaultConfig(config.CacheExpTimeInSeconds)); err != nil {
		log.Fatal().Err(err).Msg("Cannot initialize the BigCache")
	}

	// Init UpdateCacheScheduler
	stopScheduler := make(chan bool, 1)
	cacheProvider.InitUpdateCacheScheduler(stopScheduler)

	//Initialise HTTP Servers
	hTTPServer := httpserver.NewServer(config, cacheProvider, mysql)

	var errChan = make(chan error, 1)

	go func() {
		log.Info().Msgf("starting server at %s", config.HTTPPort)
		errChan <- 	hTTPServer.Start()
	}()

	//Handle Graceful Termination
	hTTPServer.Shutdown(errChan)
}
