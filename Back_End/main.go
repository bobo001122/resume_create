package main

import (
	"Back_End/config"
	"Back_End/db"
	"Back_End/db/model"
	"Back_End/router"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	db.InitDB()
	model.InitModel()
	if !config.Config.Dev {
		gin.SetMode(gin.ReleaseMode)
	}
	router.InitRouter()
	router.Router.Run(":" + config.Config.Server.Port)
}
