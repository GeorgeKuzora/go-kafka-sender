package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

type ConfigMap struct {
	Host string
	Topic string
}

type KafkaProducer struct {
	configMap ConfigMap
	broker *kafka.Writer
	ctx context.Context
}

func (kp *KafkaProducer) NewProducer(configMap ConfigMap) KafkaProducer {
	config := kafka.WriterConfig {
		Brokers: []string{configMap.Host},
		Topic: configMap.Topic,
	}
	ctx := context.Background()
	writer := *kafka.NewWriter(config)
	return KafkaProducer {
		configMap: configMap,
		broker: &writer,
		ctx: ctx,
	}
}

func (kp *KafkaProducer) Send(message string) error {
	if kp == nil {
		return fmt.Errorf("need initialized Kafka producer")
	}
	err := kp.broker.WriteMessages(kp.ctx, kafka.Message{Value: []byte(message)})
	if err != nil {
		return fmt.Errorf("failed to send a message: %s, %w", message, err)
	}
	return nil
}

func (kp *KafkaProducer) Close() error {
	if kp == nil {
		return fmt.Errorf("need initialized Kafka producer to close")
	}
	return kp.broker.Close()
}
