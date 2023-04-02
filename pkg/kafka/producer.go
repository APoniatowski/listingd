package kafka

import (
	"log"

	"github.com/Shopify/sarama"
)

type Producer struct {
	producer sarama.SyncProducer
}

func NewProducer(brokers []string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &Producer{
		producer: producer,
	}, nil
}

func (p *Producer) Produce(topic string, key string, value []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(value),
	}

	_, _, err := p.producer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to produce message: %v", err)
		return err
	}

	return nil
}

func (p *Producer) Close() error {
	return p.producer.Close()
}
