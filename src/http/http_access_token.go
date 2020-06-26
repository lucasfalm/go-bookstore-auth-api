package http

import (
	"net/http"

	"github.com/flucas97/bookstore/auth-api/src/domain/access_token"
	"github.com/flucas97/bookstore/auth-api/src/service"
	"github.com/flucas97/bookstore/auth-api/src/utils/errors_utils"
	"github.com/gin-gonic/gin"
)

// Access point to controller
type AccessTokenHandlerInterface interface {
	GetById(*gin.Context)
	Create(c *gin.Context)
}

// only related with services
type accessTokenHandler struct {
	service service.ServiceInterface
}

// Return the access point for controller, and passing wich service it will handle
func NewAccessTokenHandler(service service.ServiceInterface) AccessTokenHandlerInterface {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := handler.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessToken

	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors_utils.NewInternalServerError("invalid json")
		c.JSON(restErr.Status, restErr)
		return
	}

	if err := handler.service.Create(at); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, at)
}
