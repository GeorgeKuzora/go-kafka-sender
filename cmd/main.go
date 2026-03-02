package main
//
// import (
// 	"github.com/GeorgeKuzora/go-kafka-sender/pkg/cli"
// )
//
// func main() {
// 	cli.Run()
// }

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Kafka broker address
	broker := "localhost:9092"
	topic := "test-topic"

	// Create Kafka writer
	writer := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	// Ensure writer is closed when done
	defer writer.Close()

	// Create message
	message := kafka.Message{
		Key:   []byte("message-key"),
		Value: []byte("Hello, Kafka!"),
		Time:  time.Now(),
	}

	// Send message
	ctx := context.Background()
	err := writer.WriteMessages(ctx, message)
	if err != nil {
		log.Fatal("Failed to write message:", err)
	}

	fmt.Println("Message sent successfully!")
}
