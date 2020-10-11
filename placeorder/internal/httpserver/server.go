package httpserver

import (
	"OMS/placeorder/config"
	"OMS/placeorder/internal/cache"
	"OMS/placeorder/internal/clients"
	"OMS/placeorder/internal/service"
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
func NewServer(config *config.Config, redis *cache.Redis) *HttpServer {
	s := Server{
		Config: config,
		Order: &service.PlaceOrder{
			RedisClient:     redis,
			InventoryClient: &clients.InventoryClient{
				Config: config,
			},
		},
	}

	// Create a HTTP server for prometheus.
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/placeOrder", func(res http.ResponseWriter, req *http.Request) {
		s.placeOrder(res, req)
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

func (s *HttpServer) Shutdown(errChan chan error, stopListener chan bool) {
	var signalChan = make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	select {
	case <-signalChan:
		stopListener <- true
		log.Info().Msg("got an interrupt, exiting...")
	case err := <-errChan:
		if err != nil {
			stopListener <- true
			log.Error().Err(err).Msg("error while running api, exiting...")
		}
	}
}