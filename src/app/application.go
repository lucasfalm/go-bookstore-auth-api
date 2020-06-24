package app

import (
	"github.com/flucas97/bookstore/auth-api/src/clients/cassandra"
	"github.com/flucas97/bookstore/auth-api/src/http"
	"github.com/flucas97/bookstore/auth-api/src/repository/db"
	"github.com/flucas97/bookstore/auth-api/src/service"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()

	repository := db.NewRepository()
	atService := service.NewService(repository)        // to use service I need to pass a repository
	atHandler := http.NewAccessTokenHandler(atService) // to use controller I need to pass a service
	// atHandler := http.NewAccessTokenHandler(service.NewService(db.NewRepository()))

	MapURL(atHandler)

	router.Run(":8080")
}
