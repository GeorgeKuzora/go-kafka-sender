package args

import (
	"errors"
	"os"

	"github.com/GeorgeKuzora/go-kafka-sender/pkg/config"
)

type Args struct {
	Url      string
	FilePath string
}

func (a *Args) ToConfig() config.Config {
	return config.Config{
		Url: a.Url,
	}
}

func GetArgs() (Args, error) {
	args := os.Args
	if len(args) < 3 {
		return Args{}, errors.New("there should be at least 2 arguments")
	}
	return Args{Url: args[1], FilePath: args[2]}, nil
}
