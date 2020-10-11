package kafka

import (
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	kafkaWriter *kafka.Writer
}
func(k *KafkaProducer) InitKafkaProducer(kafkaBrokerUrls []string, clientId string, topic string) (w *kafka.Writer, err error) {
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: clientId,
	}

	config := kafka.WriterConfig{
		Brokers:          kafkaBrokerUrls,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
	}
	w = kafka.NewWriter(config)
	k.kafkaWriter = w
	return w, nil
}