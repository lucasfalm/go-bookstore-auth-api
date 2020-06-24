package app

import "github.com/flucas97/bookstore/auth-api/src/http"

func MapURL(handler http.AccessTokenHandlerInterface) { 
	router.GET("oauth/access_token/:access_token_id", handler.GetById)
	router.POST("oauth/access_token", handler.Create)
}