package args

import (
	"errors"
	"os"

	"github.com/GeorgeKuzora/go-kafka-sender/pkg/config"
)

type Args struct {
	Url      string
	FilePath string
	Topic    string
}

func (a *Args) ToConfig() config.Config {
	return config.Config{
		Url: a.Url,
		Topic: a.Topic,
	}
}

func GetArgs() (Args, error) {
	args := os.Args
	if len(args) < 4 {
		return Args{}, errors.New("there should be at least 3 arguments")
	}
	return Args{Url: args[1], FilePath: args[2], Topic: args[3]}, nil
}
