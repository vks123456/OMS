package httpserver

import (
	"OMS/inventory/config"
	"OMS/inventory/internal/api/request"
	"OMS/inventory/internal/service"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Server struct {
	Config    *config.Config
	Inventory *service.Inventory
}

func (s *Server) getQuantity(res http.ResponseWriter, req *http.Request) {
	getQReq := &request.GetQuantity{}

	if err := json.NewDecoder(req.Body).Decode(getQReq); err != nil {
		log.Error().Msgf("Exception occurred while decoding option request - Body: %v, Error: %v", req.Body, err)
		res.Write([]byte(err.Error()))
		return
	}
	log.Info().Msgf("Got Get Quantity Request : %+v", getQReq)

	invRes, err := s.Inventory.GetProductQuantity(getQReq)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	r, _ := json.Marshal(invRes)
	res.Write(r)
}

func (s *Server) addQuantity(res http.ResponseWriter, req *http.Request) {
	addQReq := &request.AddQuantity{}

	if err := json.NewDecoder(req.Body).Decode(addQReq); err != nil {
		log.Error().Msgf("Exception occurred while decoding option request - Body: %v, Error: %v", req.Body, err)
		res.Write([]byte(err.Error()))
		return
	}
	log.Info().Msgf("Got Add Quantity Request : %+v", addQReq)

	if err := s.Inventory.AddProductQuantity(addQReq); err != nil {
		res.Write([]byte(err.Error()))
	}
}

func (s *Server) addNegativeQuantity(res http.ResponseWriter, req *http.Request) {
	addQReq := &request.AddQuantity{}

	if err := json.NewDecoder(req.Body).Decode(addQReq); err != nil {
		log.Error().Msgf("Exception occurred while decoding option request - Body: %v, Error: %v", req.Body, err)
		res.Write([]byte(err.Error()))
		return
	}
	log.Info().Msgf("Got Add Negative Quantity Request : %+v", addQReq)

	if err := s.Inventory.AddNegativeProductQuantity(addQReq); err != nil {
		res.Write([]byte(err.Error()))
	}
}
