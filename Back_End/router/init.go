// do not modify
package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"Back_End/config"
)

var Router *gin.Engine

func InitRouter() {
	Router = gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	if config.Config.Dev {
		corsConfig.AllowAllOrigins = true
	} else {
		corsConfig.AllowOrigins = config.Config.Server.AllowOrigins
	}
	Router.Use(cors.New(corsConfig))
	SetRouter()
}
