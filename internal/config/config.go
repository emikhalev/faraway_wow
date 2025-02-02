package config

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sync"

	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	Server Server `yaml:"server"`
	Logger Logger `yaml:"logger"`
	PoW    PoW    `yaml:"pow"`
}

//nolint:gochecknoglobals
var (
	configFilePath = "/.env"

	config Config
	once   sync.Once
)

//nolint:gochecknoinits
func init() {
	flag.StringVar(&configFilePath, "config", "", "path to config file")

	flag.Parse()
}

func Get(ctx context.Context) Config {
	once.Do(func() {
		file, err := os.ReadFile(configFilePath)
		if err != nil {
			panic(fmt.Sprintf("cannot read config file: %v", err))
			return
		}

		if err := yaml.Unmarshal(file, &config); err != nil {
			panic(fmt.Sprintf("cannot read config file: %v", err))
		}
	})

	return config
}
