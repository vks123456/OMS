package httpserver

import (
	"OMS/inventory/config"
	"OMS/inventory/internal/cache"
	"OMS/inventory/internal/db"
	"OMS/inventory/internal/service"
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
func NewServer(config *config.Config, appCache *cache.AppCache, mySql *db.MysqlDB) *HttpServer {
	s := Server{
		Config: config,
		Inventory: &service.Inventory{
			AppCache: appCache,
			Mysql:    mySql,
		},
	}

	// Create a HTTP server for prometheus.
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/getQuantity", func(res http.ResponseWriter, req *http.Request) {
		s.getQuantity(res, req)
	})
	serveMux.HandleFunc("/addQuantity", func(res http.ResponseWriter, req *http.Request) {
		s.addQuantity(res, req)
	})
	serveMux.HandleFunc("/addNegativeQuantity", func(res http.ResponseWriter, req *http.Request) {
		s.addNegativeQuantity(res, req)
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
