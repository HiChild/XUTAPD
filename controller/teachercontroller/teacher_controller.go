package teachercontroller

import (
	"XUTAPD/common"
	"XUTAPD/models"
	"XUTAPD/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	//绑定前端参数
	var teacher models.Teacher
	ctx.Bind(&teacher)

	//数据验证
	teacherName := teacher.TeacherName
	password := teacher.Password

	if len(teacherName) == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户名不能为空")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}
	if isTeacherNameExists(DB, teacher.TeacherName) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户名已存在")
		return
	}

	//密码加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "加密失败")
		return
	}

	//保存数据库
	teacher.Password = string(hashPassword)
	if err := DB.Create(&teacher).Error; err != nil {
		response.Fail(ctx, nil, "数据库创建失败")
		return
	}

	response.Success(ctx, gin.H{"teacher": teacher}, "注册成功")
}

func isTeacherNameExists(DB *gorm.DB, teacherName string) bool {
	var teacher models.Teacher
	DB.Where("teacher_name = ?", teacherName).First(&teacher)
	if teacher.ID != 0 {
		return true
	}
	return false
}