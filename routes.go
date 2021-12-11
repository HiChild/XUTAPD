package main

import (
	"XUTAPD/controller"
	"XUTAPD/controller/studentcontroller"
	"XUTAPD/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	authGroup := r.Group("auth")
	authGroup.POST("/register", controller.Register)
	authGroup.POST("/login", controller.Login)
	authGroup.GET("/info", middleware.AuthMiddleWare(),controller.GetInfo)

	studentGroup := r.Group("student")
	studentGroup.POST("/register", studentcontroller.Register)
	studentGroup.POST("/login", studentcontroller.Login)
	studentGroup.GET("/info", middleware.AuthMiddleWare(),studentcontroller.GetInfo)

	return r
}