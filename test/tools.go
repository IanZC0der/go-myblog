package test

import (
	"github.com/IanZC0der/go-myblog/conf"
	"github.com/IanZC0der/go-myblog/ioc"

	_ "github.com/IanZC0der/go-myblog/apps"
)

// configs and utils for testing

func DevelopmentSetup() {
	// err := conf.LoadConfigFromToml("./test/config.toml")
	err := conf.LoadConfigFromEnv()

	if err != nil {
		panic(err)
	}

	if err := ioc.DefaultControllerContainer().Init(); err != nil {
		panic(err)
	}
}
