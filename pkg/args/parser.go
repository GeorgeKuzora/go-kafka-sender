package args

import (
	"errors"
	"flag"
)

type Args struct {
	Url      string
	FilePath string
}

func GetArgs() (Args, error) {
	args := flag.Args()
	if len(args) < 3 {
		return Args{}, errors.New("there should be at least 2 arguments")
	}
	return Args{Url: args[1], FilePath: args[2]}, nil
}
