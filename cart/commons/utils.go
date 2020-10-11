package commons

import (
	"OMS/cart/internal/api/request"
	"errors"
)

func ValidateAddToCartReq(req *request.Cart) error {
	if req == nil {
		return errors.New("INVALID GET QUANTITY REQUEST")
	}
	if req.OrderId == "" {
		return errors.New("INVALID ORDER ID")
	}
	if req.User == nil || req.User.UserId == "" || req.User.Address == "" {
		return errors.New("INVALID USER INFO")
	}
	if req.ItemInfo == nil {
		return errors.New("INVALID ITEM INFO")
	}

	return nil
}
