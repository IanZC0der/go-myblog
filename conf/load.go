package conf

import "github.com/BurntSushi/toml"

var (
	config *Config
)

func C() *Config {
	if config == nil {

		panic("load the config first")

	}
	return config
}

func LoadConfigFromToml(filepath string) error {

	conf := DefaultConfig()
	_, err := toml.DecodeFile(filepath, conf)

	if err != nil {
		return err
	}

	return nil
}
