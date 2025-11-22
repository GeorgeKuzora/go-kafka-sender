package cli

import (
	"fmt"
	"os"

	"github.com/GeorgeKuzora/go-kafka-sender/pkg/args"
	"github.com/GeorgeKuzora/go-kafka-sender/pkg/config"
	"github.com/GeorgeKuzora/go-kafka-sender/pkg/fs"
)

func Run() {
	args, err := args.GetArgs()
	if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}	
	config := config.Configure(args.ToConfig())
	fmt.Println(config)

	file, err := os.Open(args.FilePath)
	if err != nil {
		fmt.Printf("failed to open a file %s", args.FilePath)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("failed to close a file %s\n", args.FilePath)
		}
	}()

	for line, err := range fs.IterateOverFile(file) {
		if err != nil {
			fmt.Printf("failed to read a file %s", args.FilePath)
			return
		}
		fmt.Println(line)
	}
}
