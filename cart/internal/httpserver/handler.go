package httpserver

import (
	"OMS/cart/internal/api/request"
	"OMS/cart/internal/service"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Server struct {
	Cart *service.CartService
}

func (s Server) addToCart(res http.ResponseWriter, req *http.Request) {
	cartReq := &request.Cart{}

	if err := json.NewDecoder(req.Body).Decode(&cartReq); err != nil {
		log.Error().Msgf("Exception occurred while decoding request - Body: %v, Error: %v", req.Body, err)
		res.Write([]byte(err.Error()))
		return
	}

	log.Info().Interface("Cart Request : ", cartReq).Msg("Got Add to Cart Req")
	cartRes, err := s.Cart.AddToCart(cartReq)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	cartResByte, _ := json.Marshal(cartRes)
	res.Write(cartResByte)
}
