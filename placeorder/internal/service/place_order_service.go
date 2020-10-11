package service

import (
	"OMS/placeorder/commons"
	"OMS/placeorder/internal/api/request"
	"OMS/placeorder/internal/api/response"
	"OMS/placeorder/internal/cache"
	"OMS/placeorder/internal/clients"
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog/log"
)

type PlaceOrder struct {
	RedisClient     *cache.Redis
	InventoryClient *clients.InventoryClient
}

func (o *PlaceOrder) PlaceOrder(req *request.PlaceOrder) (*response.PlaceOrder, error) {
	if err := commons.ValidatePlaceOrderReq(req); err != nil {
		return nil, err
	}

	o.toPay(req)

	//Place order
	status, err := o.RedisClient.Get(req.OrderId)
	if err != nil {
		log.Error().Err(err).Msg("Unable to fetch status from redis")
		return nil, err
	}
	switch status {
	case "FAILED":
		return nil, errors.New("FAILED ORDER")
	case "PAYMENT_FAILED":
	case "BLOCKED":
		if err := o.addItems(req); err != nil { // Add blocked or failed payment items back to Inventory
			return nil, err
		}
	case "PAYMENT_SUCCESS": // Check and update inventory for negative stock
		if err := o.checkAndUpdateNegativeInventory(req); err != nil {
			o.RedisClient.Set(req.OrderId, "REFUND", 24*time.Hour)
			return nil, err
		}
	}
	return &response.PlaceOrder{
		OrderId: req.OrderId,
		Status:  "SUCCESS",
	}, nil
}

/*
This method will check blocked order status for requested orderId every 2 second for 3 minutes.
If blocked order found, then initiate the payment
*/
func (o *PlaceOrder) toPay(req *request.PlaceOrder) {
	toPayScheduler := time.NewTicker(2 * time.Second)
	// Create a context that will be canceled in 3 mins.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			toPayScheduler.Stop()
			return
		case <-toPayScheduler.C:
			if v, _ := o.RedisClient.Get(req.OrderId); v == "FAILED" {
				toPayScheduler.Stop()
				return
			}
			if v, _ := o.RedisClient.Get(req.OrderId); v == "BLOCKED" {
				toPayScheduler.Stop()
				if err := initPayment(); err != nil {
					o.RedisClient.Set(req.OrderId, "PAYMENT_FAILED", 30*time.Minute)
					return
				}
				o.RedisClient.Set(req.OrderId, "PAYMENT_SUCCESS", 30*time.Minute)
				return
			}
		}
	}

}

/*
This method will call payment gateway with 1 minute timeout
*/
func initPayment() error {
	// here call payment gateway with 1 minute timeout
	return nil
}

func (o *PlaceOrder) addItems(req *request.PlaceOrder) error {
	addItemsReq := &request.AddQuantity{
		ProductId:   req.ProductId,
		AddQuantity: req.Quantity,
	}
	return o.InventoryClient.AddItemsToInventory(addItemsReq)
}

func (o *PlaceOrder) checkAndUpdateNegativeInventory(req *request.PlaceOrder) error {
	addItemsReq := &request.AddQuantity{
		ProductId:   req.ProductId,
		AddQuantity: req.Quantity,
	}
	return o.InventoryClient.CheckAndUpdateNegativeInventory(addItemsReq)
}
