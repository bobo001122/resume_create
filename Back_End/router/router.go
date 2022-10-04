package router

import (
	"Back_End/controller"
	"Back_End/utility/middleware"
)

func SetRouter() {
	Router.POST("/api/user/register", controller.Register)
	Router.POST("/api/user/login", controller.Login)
	Router.GET("/api/user", middleware.Authorization, controller.GetUserInfo)

	// modifyline 14-17
	Router.POST("/api/resume", middleware.Authorization, controller.AddResume)
	Router.DELETE("/api/resume", middleware.Authorization, controller.DeleteResume)
	Router.PUT("/api/resume", middleware.Authorization, controller.UpdateResume)
	Router.GET("/api/resume", middleware.Authorization, controller.GetResumes)
}
