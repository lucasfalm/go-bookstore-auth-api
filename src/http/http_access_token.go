package http

import (
	"net/http"

	"github.com/flucas97/bookstore/auth-api/src/domain/access_token"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.ServiceInterface
}

func NewAccessTokenHandler(service access_token.ServiceInterface) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, _ := handler.service.GetById(c.Param("access_token_id"))
	c.JSON(http.StatusOK, accessToken)
}
