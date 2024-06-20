package middlewares

import (
	// "context"

	"github.com/IanZC0der/go-myblog/apps/token"
	"github.com/IanZC0der/go-myblog/ioc"
	"github.com/gin-gonic/gin"
)

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{
		tkSvc: ioc.DefaultControllerContainer().Get(token.AppName).(token.Service),
	}
}

type AuthMiddleware struct {
	tkSvc token.Service
}

func (a *AuthMiddleware) Auth(c *gin.Context) {

	// req := token.NewValidateToken("")
	// return a.tkSvc.ValidateToken(ctx, req)
}
