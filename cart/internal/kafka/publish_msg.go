package kafka

import (
	"context"

	"time"

	"github.com/segmentio/kafka-go"
)

func (k *KafkaProducer) PushMsg(parent context.Context, key, value []byte) (err error) {
	message := kafka.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}

	return k.kafkaWriter.WriteMessages(parent, message)
}
