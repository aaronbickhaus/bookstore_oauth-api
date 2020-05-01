package app

import (
	"github.com/aaronbickhaus/bookstore_oauth-api/src/domain/access_token"
	"github.com/aaronbickhaus/bookstore_oauth-api/src/http/http_access_token"
	"github.com/aaronbickhaus/bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)
var (
	router = gin.Default()
)

func StartApplication() {

	atHandler := http_access_token.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.PUT("/oauth/access_token", atHandler.UpdateExpirationTime)
	_ = router.Run(":8080")
}

