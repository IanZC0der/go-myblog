package api

import (
	"net/http"

	"github.com/IanZC0der/go-myblog/apps/token"
	"github.com/gin-gonic/gin"
)

type TokenApiHandler struct {
	svc token.Service
}

func NewTokenApiHandler(tokenServiceImpl token.Service) *TokenApiHandler {
	return &TokenApiHandler{
		svc: tokenServiceImpl,
	}
}

func (h *TokenApiHandler) Registry(router gin.IRouter) {
	v1 := router.Group("v1")
	v1.POST("/tokens/", h.Login)
	v1.DELETE("/tokens/", h.Logout)
}

func (h *TokenApiHandler) Login(c *gin.Context) {
	// get req params

	newReq := token.NewLoginRequest()
	err := c.BindJSON(newReq)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// login
	tk, err := h.svc.Login(c.Request.Context(), newReq)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	//return response
	c.JSON(http.StatusOK, tk)

}

func (h *TokenApiHandler) Logout(*gin.Context) {

}
