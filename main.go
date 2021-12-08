package main

import (
	"XUTAPD/common"
	"github.com/gin-gonic/gin"
)


func main() {
	common.InitDB()

	r := gin.Default()
	r = CollectRoutes(r)

	port := "8080"
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}




