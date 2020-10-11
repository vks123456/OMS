package kafka

import (
	"OMS/placeorder/internal/cache"
	"OMS/placeorder/internal/clients"
	"OMS/placeorder/internal/service"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	Reader          *kafka.Reader
	OrderService    *service.PlaceOrder
	InventoryClient *clients.InventoryClient
	RedisClient     *cache.Redis
}

func (k *KafkaConsumer) InitKafkaConsumer(brokers []string, clientId string, topic string) *kafka.Reader {
	// make a new reader that consumes from topic-A
	config := kafka.ReaderConfig{
		Brokers:         brokers,
		GroupID:         clientId,
		Topic:           topic,
		MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: -1,
	}

	reader := kafka.NewReader(config)
	k.Reader = reader
	return reader
}
