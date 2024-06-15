package test

import "github.com/IanZC0der/go-myblog/conf"

// configs and utils for testing

func DevelopmentSetup() {
	// err := conf.LoadConfigFromToml("./test/config.toml")
	err := conf.LoadConfigFromEnv()

	if err != nil {
		panic(err)
	}
}
