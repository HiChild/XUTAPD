package studentcontroller

import (
	"XUTAPD/common"
	"XUTAPD/dto"
	"XUTAPD/models"
	"XUTAPD/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	var student models.Student
	//绑定参数
	ctx.Bind(&student)
	studentName := student.StudentName
	password := student.Password
	if len(studentName) == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户名不能为空")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}
	if isStudentNameExists(DB, student.StudentName) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户名已存在")
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "加密失败")
		return
	}

	newStudent := models.Student{
		StudentName: student.StudentName,
		Password:    string(hashPassword),
	}

	if err := DB.Create(&newStudent).Error; err != nil {
		response.Fail(ctx, gin.H{"err": err}, "数据库创建失败")
		return
	}

	response.Success(ctx, gin.H{"student": newStudent}, "注册成功")
}

func Login(ctx *gin.Context) {
	DB := common.GetDB()
	//绑定参数
	var student models.Student
	ctx.Bind(&student)
	studentName := student.StudentName
	password := student.Password

	//数据验证
	if len(studentName) == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户名不能为空")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}
	//判断用户名是否存在,复用之前的user变量
	DB.Where("student_name = ?", studentName).First(&student)
	if student.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(password)); err != nil {
		response.Fail(ctx, nil, "密码错误")
		return
	}
	//发放token
	token, err := common.ReleaseTokenStudent(student)
	if err != nil {
		response.Fail(ctx, nil, "系统异常")
		log.Printf("token generator error : %v", err)
		return
	}
	//返回结果
	response.Success(ctx, gin.H{"token": token}, "登录成功")
}

func isStudentNameExists(DB *gorm.DB, studentName string) bool {
	var student models.Student
	DB.Where("student_name = ?", studentName).First(&student)
	if student.ID != 0 {
		return true
	}
	return false
}

func GetInfo(ctx *gin.Context) {
	student, _ := ctx.Get("student")

	response.Success(ctx, gin.H{"student": dto.ToStudentDTO(student.(models.Student))}, "ok")
}
