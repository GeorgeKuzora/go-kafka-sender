package cli

import (
	"fmt"
	"os"

	"github.com/GeorgeKuzora/go-kafka-sender/pkg/args"
	"github.com/GeorgeKuzora/go-kafka-sender/pkg/config"
)

func Run() {
	args, err := args.GetArgs()
	if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}	
	fmt.Println(args)
	config := config.Configure(args.ToConfig())
	fmt.Println(config)
}
