package test

import "github.com/IanZC0der/go-myblog/conf"

// configs and utils for testing

func DevelopmentSetup() {
	err := conf.LoadConfigFromToml("./config.toml")

	if err != nil {
		panic(err)
	}
}
