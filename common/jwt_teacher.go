package common

import (
	"XUTAPD/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//发放Token 给User
func ReleaseTokenTeacher(teacher models.Teacher) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	//设置Claims
	studentClaims := &Claims{
		UserId:         teacher.ID,
		StandardClaims: jwt.StandardClaims{
			//过期时间
			ExpiresAt: expirationTime.Unix(),
			//发放时间
			IssuedAt : time.Now().Unix(),
			//发放机构
			Issuer: "xutapd.dev",
			//主题
			Subject: "teacher Token",
		},
	}

	//选择HS的加密算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, studentClaims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//解析Token
func ParseTokenTeacher(tokenString string) (*jwt.Token, *Claims, error){
	teacherClaims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, teacherClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, teacherClaims, err
}

