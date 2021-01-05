package app

import (
	"github.com/Ferza17/Oauth-APIsrc/clients/cassandra"
	"github.com/Ferza17/Oauth-API/src/http"
	"github.com/Ferza17/Oauth-API/src/repository/db"
	"github.com/Ferza17/Oauth-API/src/repository/rest"
	access_token_service "github.com/Ferza17/Oauth-API/src/services/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router =  gin.Default()
)

func StartApplication()  {
	session := cassandra.GetSession()
	defer session.Close()

	atHandler := http.NewAccessTokenHandler(
		access_token_service.NewService(rest.NewRepository(), db.NewRepository()))

	// GET Access Token
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	// POST (Create) AccessToken
	router.POST("/oauth/access_token/", atHandler.Create)

	// PUT (Update) AccessToken
	router.PUT("/oauth/access_token/:access_token_id")

	if err := router.Run(":8080"); err != nil {
		panic("Cant connect to the App")
	}

}