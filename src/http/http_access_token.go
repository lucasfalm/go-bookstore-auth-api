package http

import (
	"net/http"

	"github.com/flucas97/bookstore/auth-api/src/domain/access_token"
	"github.com/gin-gonic/gin"
)

var (
	AccessTokenHandler HandlerInterface = &accessTokenHandler{}
)

type HandlerInterface interface {
	GetById(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) *accessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "implement me")
}
