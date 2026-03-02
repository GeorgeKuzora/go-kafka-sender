package config

import (
	"github.com/GeorgeKuzora/go-kafka-sender/pkg/args"
	"github.com/GeorgeKuzora/go-kafka-sender/pkg/kafka"
)

type Config struct {
	Url string `yaml:"url"`
	Topic string `yaml:"url"`
}

func (c *Config) merge(oc Config) *Config {
	if oc.Url != "" {
		c.Url = oc.Url
	}
	if oc.Topic != "" {
		c.Topic = oc.Topic
	}
	return c
}

func (c *Config) ToConfigMap() kafka.ConfigMap {
	return kafka.ConfigMap {
		Host: c.Url,
		Topic: c.Topic,
	}
}

func FromArgs(args args.Args) Config {
	return Config{
		Url: args.Url,
		Topic: args.Topic,
	}
}


func Configure(argsConfig Config) Config {
	config := createBaseConfig()
	config = *config.merge(argsConfig)
	return config
}

func createBaseConfig() Config {
	return Config{
		Url: "http://127.0.0.1:9092",
		Topic: "",
	}
}
