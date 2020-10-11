package kafka

import (
	"OMS/placeorder/internal/api/request"
	"context"
	"encoding/json"
	"time"

	"github.com/rs/zerolog/log"
	_ "github.com/segmentio/kafka-go/snappy"
)

func (k *KafkaConsumer) ListenMsg(stopListener chan bool) {
	log.Info().Msg("Going to listen msg...")

	for {
		select {
		case <-stopListener:
			return
		default:
			m, err := k.Reader.ReadMessage(context.Background())
			if err != nil {
				log.Error().Msgf("error while receiving message: %s", err.Error())
				continue
			}

			if err != nil {
				log.Error().Msgf("error while receiving message: %s", err.Error())
				continue
			}
			log.Info().Msgf("message at topic/partition/offset %v/%v/%v: %s\n", m.Topic, m.Partition, m.Offset, string(m.Value))

			//Check & Block Inventory stock
			k.checkAndBlockInventory(m.Value)

		}
	}
}

func (k *KafkaConsumer) checkAndBlockInventory(msg []byte) {
	cartReq := &request.Cart{}
	if err := json.Unmarshal(msg, cartReq); err != nil {
		log.Error().Err(err).Msg("Error while unmarshalling cart req")
		return
	}
	inventoryReq := &request.GetQuantity{
		ProductId:     cartReq.ItemInfo.ItemId,
		FromCache:     false,
		BlockQuantity: cartReq.ItemInfo.Quantity,
	}

	if _, err := k.InventoryClient.CheckAndBlockInventory(inventoryReq); err != nil {
		log.Error().Err(err).Msg("unable to block inventory")
		k.RedisClient.Set(cartReq.OrderId, "FAILED", 5*time.Minute)
		return
	}

	k.RedisClient.Set(cartReq.OrderId, "BLOCKED", 5*time.Minute)
}
