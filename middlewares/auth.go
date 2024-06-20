package middlewares

import (
	// "context"

	"net/http"

	"github.com/IanZC0der/go-myblog/apps/token"
	"github.com/IanZC0der/go-myblog/exception"
	"github.com/IanZC0der/go-myblog/ioc"
	"github.com/IanZC0der/go-myblog/response"
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

	// get the token from the token
	accessToken, err := c.Cookie(token.TOKEN_COOKIE_NAME)

	if err != nil {
		if err == http.ErrNoCookie {
			response.Failed(c, exception.NewAuthFailed("cookie %s not found", token.TOKEN_GIN_KEY_IN_CONTEXT))
		}
		response.Failed(c, err)
		return
	}

	// create the request
	req := token.NewValidateToken(accessToken)
	// validate the token
	tk, err := a.tkSvc.ValidateToken(c.Request.Context(), req)

	if err != nil {
		response.Failed(c, err)
		return
	}

	// initialize c.keys if nil
	if c.Keys == nil {
		c.Keys = map[string]any{}
	}

	// put token in the context
	c.Keys[token.TOKEN_GIN_KEY_IN_CONTEXT] = tk
}
