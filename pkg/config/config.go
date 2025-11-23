package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/GeorgeKuzora/go-kafka-sender/pkg/args"
	"github.com/GeorgeKuzora/go-kafka-sender/pkg/kafka"
	"gopkg.in/yaml.v3"
)

type filePathStatus int

const (
	valid filePathStatus = iota
	notExits
	permissionDenied
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
	config = *config.merge(
		getSystemConfig(),
	).merge(
		getUserConfig(),
	).merge(
		argsConfig,
	)
	return config
}

func createBaseConfig() Config {
	return Config{
		Url: "http://127.0.0.1:9092",
		Topic: "",
	}
}

func getUserConfig() Config {
	path, err := findUserConfigFilePath()
	if err != nil {
		return Config{}
	}
	config, err := readConfigFile(path)
	if err != nil {
		return Config{}
	}
	return config
}

func getSystemConfig() Config {
	return Config{}
}

func findUserConfigFilePath() (string, error) {
	if configPath := os.Getenv("GO_KAFKA_CONFIG_PATH"); configPath != "" {
		return configPath, nil
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}
	return filepath.Join(homeDir, ".config", "gokafka", "config.yaml"), nil
}

func readConfigFile(filePath string) (Config, error) {
	switch validateFilePath(filePath) {
	case notExits:
		return Config{}, nil
	case permissionDenied:
		return Config{}, fmt.Errorf("can't read config file %s", filePath)
	case valid:
		config := &Config{}
		err := loadYaml(filePath, config)
		if err != nil {
			return Config{}, err
		}
		return *config, nil
	}
	return Config{}, fmt.Errorf("unknown config file status %s", filePath)
}

func loadYaml(path string, config *Config) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("can't read config %w", err)	
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return fmt.Errorf("can't parse YAML config %w", err)	
	}
	return nil
}

func validateFilePath(filePath string) filePathStatus {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return notExits
	} else if err != nil {
		return permissionDenied
	} else {
		return valid
	}
}
