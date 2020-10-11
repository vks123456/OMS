package main

import (
	cfg "OMS/placeorder/config"
	"OMS/placeorder/internal/cache"
	"OMS/placeorder/internal/clients"
	"OMS/placeorder/internal/httpserver"
	"OMS/placeorder/internal/kafka"
	"OMS/placeorder/internal/service"
	"strings"

	"github.com/rs/zerolog/log"
)

func main() {
	//Initialise Configs
	config := cfg.InitConfigStore()
	log.Info().Msg("Starting place order service...")

	// Init Redis
	log.Info().Msg("Initializing Redis...")
	r := &cache.Redis{}
	r.NewSession(config)

	//Initialise kafka consumer
	log.Info().Msg("Initializing Kafka consumer...")
	k := kafka.KafkaConsumer{
		OrderService: &service.PlaceOrder{},
		InventoryClient: &clients.InventoryClient{
			Config: config,
		},
		RedisClient: r,
	}
	consumer := k.InitKafkaConsumer(strings.Split(config.KafkaBrokerUrl, ","), config.KafkaClientId, config.KafkaTopic)
	defer consumer.Close()

	//Listen Msg
	stopListener := make(chan bool, 1)
	go k.ListenMsg(stopListener)

	//Initialise HTTP Servers
	hTTPServer := httpserver.NewServer(config, r)

	var errChan = make(chan error, 1)

	go func() {
		log.Info().Msgf("starting server at %s", config.HTTPPort)
		errChan <- 	hTTPServer.Start()
	}()

	//Handle Graceful Termination
	hTTPServer.Shutdown(errChan, stopListener)
}
