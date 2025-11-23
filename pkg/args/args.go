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
	var filepath, topic, url string
	args_len := len(args)
	if args_len < 2 {
		return Args{}, errors.New("there should be at least 1 argument")
	} else if args_len < 3 {
		filepath = args[1]
		topic = ""
		url = ""
	} else if args_len < 4 {
		filepath = args[2]
		topic = args[1]
		url = ""
	} else {
		filepath = args[3]
		topic = args[2]
		url = args[1]
	}
	return Args{Url: url, FilePath: filepath, Topic: topic}, nil
}
