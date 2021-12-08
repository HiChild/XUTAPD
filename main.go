package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type User struct {
	gorm.Model
	UserName     string `form:"username" gorm:"type:varchar(20);not null;unique"`
	Password string `form:"password" gorm:"size:255"`
}

func main() {

	//需要全局db
	dsn := "root:7WzhXRTJdSEZWknE@tcp(120.53.228.79:3306)/xutapd?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("connect Mysql Error: %v", err)
	}
	DB.AutoMigrate(&User{}) //自动迁移

	//gin 框架
	r := gin.Default()
	accountGroup := r.Group("account")
	{
		//注册
		accountGroup.POST("/register", func(ctx *gin.Context) {
			var user User
			//绑定参数
			if err = ctx.Bind(&user); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"code":422,
					"msg":"参数绑定失败",
				})
				return
			}
			if isUserNameExists(DB, user.UserName) {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"code": 422,
					"msg": "用户已存在",
				})
				return
			}
			if err = DB.Create(&user).Error; err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"code": 422,
					"msg": "注册失败",
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"code": 200,
					"msg":"注册成功",
					"data":user,
				})
			}
		})

		//登录
		accountGroup.POST("/login", func(ctx *gin.Context) {
		})
	}

	r.Run(":8080")
}

func isUserNameExists(DB *gorm.DB, username string) bool {
	var user User
	DB.Where("user_name = ?", username).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}


