package middleware

import (
	"XUTAPD/common"
	"XUTAPD/models"
	"XUTAPD/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleWareStudent() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization Header
		tokenString := ctx.GetHeader("Authorization")

		//validate token format
		//token格式不正确
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			response.Fail(ctx, nil, "权限不足")
			//使用本函数抛弃本次请求
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
			ctx.Abort()
			return
		}

		//通过后获取userId,
		userId := claims.UserId

		DB := common.GetDB()
		var student models.Student
		DB.First(&student, userId)

		//查出用户
		if student.ID == 0 {
			response.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
			ctx.Abort()
			return
		}

		//加入到上下文中
		ctx.Set("student", student)

		//继续向下执行!!!!!!!!
		ctx.Next()
	}
}