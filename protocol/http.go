package protocol

import (
	"context"
	"fmt"
	"net/http"

	"time"

	"github.com/IanZC0der/go-myblog/conf"
	"github.com/IanZC0der/go-myblog/ioc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewHttpServer() *HttpServer {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	ioc.DefaultApiHandlerContainer().RouterRegistry(r.Group("/api/myblog"))
	return &HttpServer{
		sver: &http.Server{
			Addr:    conf.C().App.HttpAddress(),
			Handler: r,
		},
	}
}

type HttpServer struct {
	sver *http.Server
}

func (s *HttpServer) Run() error {
	fmt.Printf("Server starts at %s\n", conf.C().App.HttpAddress())
	return s.sver.ListenAndServe()
}

func (s *HttpServer) Close(ctx context.Context) {
	s.sver.Shutdown(ctx)
}
