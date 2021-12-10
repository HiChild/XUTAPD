package main

import (
	"XUTAPD/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)


func main() {
	//必须先初始化viper配置
	initConfig()
	common.InitDB()
	r := gin.Default()
	r = CollectRoutes(r)

	//使用viperGetXXX读取配置文件
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func initConfig()  {
	//获取当前工作目录
	workDir, _ := os.Getwd()
	//设置配置文件的名称
	viper.SetConfigName("application")
	//设置文件类型
	viper.SetConfigType("yml")

	viper.AddConfigPath(workDir + "/config")
	//或者
	//viper.AddConfigPath("./config")

	//读取文件
	err := viper.ReadInConfig()
	if err != nil {
		panic("")
	}
}

