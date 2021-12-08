package main

import (
	"XUTAPD/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	accountGroup := r.Group("auth")
	accountGroup.POST("/register", controller.Register)
	accountGroup.POST("/login", func(ctx *gin.Context) {
	})


	return r
}