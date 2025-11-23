package args

import (
	"errors"
	"os"
)

type Args struct {
	Url      string
	FilePath string
	Topic    string
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
