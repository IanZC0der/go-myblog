package main

import (
	"fmt"
	"os"

	tokenAPIHandler "github.com/IanZC0der/go-myblog/apps/token/api"
	tokenImpl "github.com/IanZC0der/go-myblog/apps/token/impl"
	userImpl "github.com/IanZC0der/go-myblog/apps/user/impl"
	"github.com/IanZC0der/go-myblog/conf"
	"github.com/gin-gonic/gin"
)

func main() {
	// load config

	err := conf.LoadConfigFromToml("etc/development.toml")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// init controller: user controller, token controller, api handler
	userSvcImpl := userImpl.NewUserServiceImpl()
	tokenSvcImpl := tokenImpl.NewTokenServiceImpl(userSvcImpl)

	tokenApiHandler := tokenAPIHandler.NewTokenApiHandler(tokenSvcImpl)

	r := gin.Default()

	tokenApiHandler.Registry(r.Group("/api/myblog"))

	// start http server, register router

	err = r.Run(conf.C().App.HttpAddress())

	if err != nil {
		fmt.Println(err)
	}

}
