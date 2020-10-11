package service

import (
	"OMS/cart/internal/api/request"
	"OMS/cart/internal/api/response"
	"OMS/cart/internal/cache"
	"OMS/cart/internal/clients"
	"OMS/cart/internal/kafka"
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/rs/zerolog/log"
)

type CartService struct {
	InventoryClient *clients.InventoryClient
	Producer        *kafka.KafkaProducer
	RedisClient *cache.Redis
}

func (c *CartService) AddToCart(req *request.Cart) (*response.Cart, error) {
	// check cached inventory stock
	if c.getAvailableStock(req.ItemInfo) < req.ItemInfo.Quantity {
		log.Error().Msgf("Item/Id - %v/%v  is out of stock", req.ItemInfo.ItemName, req.ItemInfo.ItemId)
		return nil, errors.New("ITEM IS OUT OF STOCK")
	}

	// To place order, Push order to kafka
	reqByte, _ := json.Marshal(req)
	if err := c.Producer.PushMsg(context.Background(), []byte(req.OrderId), reqByte); err != nil {
		log.Error().Err(err).Msgf("Unable to push msg into kafka for orderId : &v", req.OrderId)
		return nil, err
	}

	//set order status against orderId into Redis
	if err := c.RedisClient.Set(req.OrderId, "INIT", 15*time.Minute); err != nil {
		log.Error().Err(err).Msgf("Unable to set order status in Redis for orderId : %v", req.OrderId)
		return nil, err
	}
	return nil, nil
}

func (c *CartService) getAvailableStock(item *request.ItemInfo) int {
	req := &request.GetQuantity{
		ProductId: item.ItemId,
		FromCache: true,
	}

	res, err := c.InventoryClient.CheckInventory(req)
	if err != nil {
		log.Error().Err(err).Msgf("Not able to check inventory of Item : %v", item.ItemId)
		return 0
	}
	if res == nil {
		log.Error().Msgf("Invalid Inventory Response for Item : %v", item.ItemId)
		return 0
	}
	return res.Quantity
}
