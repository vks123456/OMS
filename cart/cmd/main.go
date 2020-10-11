package main

import (
	cfg "OMS/cart/config"
	"OMS/cart/internal/cache"
	"OMS/cart/internal/httpserver"
	"OMS/cart/internal/kafka"
	"strings"

	"github.com/rs/zerolog/log"
)

func main() {
	//Initialise Configs
	config := cfg.InitConfigStore()
	log.Info().Msg("Starting cart service...")

	//Init Kafka
	log.Info().Msg("Initializing kafka...")
	kc := kafka.KafkaProducer{}
	kafkaProducer, err := kc.InitKafkaProducer(strings.Split(config.KafkaBrokerUrl, ","), config.KafkaClientId, config.KafkaTopic)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to initialize kafka")
	}
	defer kafkaProducer.Close()

	// Init Redis
	r := cache.Redis{}
	r.NewSession(config)

	//Initialise HTTP Servers
	hTTPServer := httpserver.NewServer(config, &kc, &r)

	var errChan = make(chan error, 1)

	go func() {
		log.Info().Msgf("starting server at %s", config.HTTPPort)
		errChan <- hTTPServer.Start()
	}()

	//Handle Graceful Termination
	hTTPServer.Shutdown(errChan)
}
