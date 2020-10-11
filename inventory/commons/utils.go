package commons

import (
	"OMS/inventory/internal/api/request"
	"errors"
)

func ValidateGetQuantityReq(req *request.GetQuantity) error {
	if req == nil {
		return errors.New("INVALID GET QUANTITY REQUEST")
	}
	if req.ProductId == "" {
		return errors.New("INVALID PRODUCT ID")
	}
	return nil
}

func ValidateAddQuantityReq(req *request.AddQuantity) error {
	if req == nil {
		return errors.New("INVALID ADD QUANTITY REQUEST")
	}
	if req.ProductId == "" {
		return errors.New("INVALID PRODUCT ID")
	}
	return nil
}
