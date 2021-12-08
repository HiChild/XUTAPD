package controller

import (
	"XUTAPD/common"
	"XUTAPD/models"
	"XUTAPD/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func Register (ctx *gin.Context) {
	DB := common.GetDB()
	var user models.User
	//绑定参数
	ctx.Bind(&user)
	username := user.UserName
	password := user.Password
	if len(username) == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户名不能为空")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}
	if isUserNameExists(DB, user.UserName) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户名已存在")
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "加密失败")
		return
	}

	newUser := models.User{
		Model:    gorm.Model{},
		UserName: username,
		Password: string(hashPassword),
	}

	DB.Create(&newUser)

	response.Success(ctx, gin.H{"user":newUser}, "注册成功")
}

func Login(ctx *gin.Context) {
	//获取参数
	DB := common.GetDB()
	var user models.User
	//绑定参数
	ctx.Bind(&user)
	username := user.UserName
	password := user.Password

	//数据验证
	if len(username) == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户名不能为空")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}
	//判断用户名是否存在,复用之前的user变量
	DB.Where("user_name = ?", username).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Fail(ctx, nil, "密码错误")
		return
	}
	//发放token
	token:= "11"
	//返回结果
	response.Success(ctx, gin.H{"token": token}, "登录成功")
}

func isUserNameExists(DB *gorm.DB, username string) bool {
	var user models.User
	DB.Where("user_name = ?", username).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}