package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

var (
	config *Config = DefaultConfig()
)

func C() *Config {
	return config
}

func LoadConfigFromToml(filepath string) error {

	// conf := DefaultConfig()
	_, err := toml.DecodeFile(filepath, config)

	if err != nil {
		return err
	}

	return nil
}

func LoadConfigFromEnv() error {
	return env.Parse(config)
}
