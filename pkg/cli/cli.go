package cli

import (
	"fmt"
	"os"

	"github.com/GeorgeKuzora/go-kafka-sender/pkg/args"
	"github.com/GeorgeKuzora/go-kafka-sender/pkg/config"
	"github.com/GeorgeKuzora/go-kafka-sender/pkg/fs"
	"github.com/GeorgeKuzora/go-kafka-sender/pkg/kafka"
)

func Run() {
	args, err := args.GetArgs()
	if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}	
	config := config.Configure(config.FromArgs(args))

	file, err := os.Open(args.FilePath)
	if err != nil {
		fmt.Printf("failed to open a file %s", args.FilePath)
		os.Exit(1)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("failed to close a file %s\n", args.FilePath)
		}
		os.Exit(1)
	}()

	for line, err := range fs.IterateOverFile(file) {
		if err != nil {
			fmt.Printf("failed to read a file %s", args.FilePath)
			return
		}
		producer := kafka.NewProducer(config.ToConfigMap())
		err := producer.Send(line)
		if err != nil {
			fmt.Printf("failed to send Kafka message")
			os.Exit(1)
		}
	}
}
