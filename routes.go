package main

import (
	"XUTAPD/controller"
	"XUTAPD/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	authGroup := r.Group("auth")
	authGroup.POST("/register", controller.Register)
	authGroup.POST("/login", controller.Login)
	authGroup.GET("/info", middleware.AuthMiddleWare(),controller.GetInfo)

	return r
}