package httpserver

import (
	"OMS/cart/config"
	"OMS/cart/internal/cache"
	"OMS/cart/internal/clients"
	"OMS/cart/internal/kafka"
	"OMS/cart/internal/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

type HttpServer struct {
	*http.ServeMux
	*config.Config
}

/*
Configure Http Server for prometheus.
*/
func NewServer(config *config.Config, producer *kafka.KafkaProducer, redis *cache.Redis) *HttpServer {

	s := Server{
		Cart: &service.CartService{
			InventoryClient: &clients.InventoryClient{
				Config: config,
			},
			Producer:    producer,
			RedisClient: redis,
		},
	}
	// Create a HTTP server for prometheus.
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/addToCart", func(res http.ResponseWriter, req *http.Request) {
		s.addToCart(res, req)
	})

	return &HttpServer{
		ServeMux: serveMux,
		Config:   config,
	}
}

/*
Start Http server.
*/
func (s *HttpServer) Start() error {

	if err := http.ListenAndServe(s.Config.Host+":"+s.Config.HTTPPort, s.ServeMux); err != nil {
		log.Error().Err(err).Msg("Unable to start http server")
		return err
	}
	return nil
}

func (s *HttpServer) Shutdown(errChan chan error) {
	var signalChan = make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	select {
	case <-signalChan:
		log.Info().Msg("got an interrupt, exiting...")
	case err := <-errChan:
		if err != nil {
			log.Error().Err(err).Msg("error while running api, exiting...")
		}
	}
}
