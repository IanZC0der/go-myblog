package protocol

import (
	"context"
	"fmt"
	"net/http"

	"github.com/IanZC0der/go-myblog/conf"
	"github.com/IanZC0der/go-myblog/ioc"
	"github.com/gin-gonic/gin"
)

func NewHttpServer() *HttpServer {
	r := gin.Default()

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
