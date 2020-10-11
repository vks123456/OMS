package httpserver

import (
	"OMS/placeorder/config"
	"OMS/placeorder/internal/api/request"
	"OMS/placeorder/internal/service"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Server struct {
	Config *config.Config
	Order  *service.PlaceOrder
}

func (s *Server) placeOrder(res http.ResponseWriter, req *http.Request) {
	orderReq := &request.PlaceOrder{}

	if err := json.NewDecoder(req.Body).Decode(orderReq); err != nil {
		log.Error().Msgf("Exception occurred while decoding option request - Body: %v, Error: %v", req.Body, err)
		res.Write([]byte(err.Error()))
		return
	}
	log.Info().Interface("PlaceOrderReq : ", orderReq).Msg("Got Place Order Request")

	r, err := s.Order.PlaceOrder(orderReq)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	rByte, _ := json.Marshal(r)
	res.Write(rByte)
}
