package api

import (
	// "net/http"

	"github.com/IanZC0der/go-myblog/apps/token"
	"github.com/IanZC0der/go-myblog/ioc"
	"github.com/IanZC0der/go-myblog/response"
	"github.com/gin-gonic/gin"
)

type TokenApiHandler struct {
	svc token.Service
}

func NewTokenApiHandler() *TokenApiHandler {
	return &TokenApiHandler{
		svc: ioc.DefaultControllerContainer().Get(token.AppName).(token.Service),
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
		// c.JSON(http.StatusBadRequest, err.Error())
		response.Failed(c, err)
		return
	}
	// login
	tk, err := h.svc.Login(c.Request.Context(), newReq)

	if err != nil {

		response.Failed(c, err)
		// c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	//return response
	response.Success(c, tk)
	// c.JSON(http.StatusOK, tk)

}

func (h *TokenApiHandler) Logout(*gin.Context) {

}
