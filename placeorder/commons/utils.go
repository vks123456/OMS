package commons

import (
	"OMS/placeorder/internal/api/request"
	"errors"
)

func ValidatePlaceOrderReq(req *request.PlaceOrder) error {
	if req == nil {
		return errors.New("INVALID PLACE ORDER REQUEST")
	}
	if req.ProductId == "" {
		return errors.New("INVALID PRODUCT ID")
	}
	if req.OrderId == "" {
		return errors.New("INVALID ORDER ID")
	}
	if req.UserId == "" {
		return errors.New("INVALID USER ID")
	}

	return nil
}
