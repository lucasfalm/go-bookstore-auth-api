package app

import (
	"github.com/flucas97/bookstore/auth-api/src/domain/access_token"
	"github.com/flucas97/bookstore/auth-api/src/http"
	"github.com/flucas97/bookstore/auth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	repository := db.NewRepository()
	atService := access_token.NewService(repository) // need to fix 'import cycle not allowed' using interfaces
	atHandler := http.NewAccessTokenHandler(atService)

	router.GET("oauth/access_token/access_token_id", atHandler.GetById)
	router.Run(":8080")
}
