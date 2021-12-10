package common

import (
	"XUTAPD/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//RSA加密的密钥
var jwtKey = []byte("a_secret_key")
//定义Claims
type Claims struct {
	UserId uint
	//嵌入jwt自定义的Claims
	jwt.StandardClaims
}

//发放Token
func ReleaseToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	//设置Claims
	claims := &Claims{
		UserId:         user.ID,
		StandardClaims: jwt.StandardClaims{
			//过期时间
			ExpiresAt: expirationTime.Unix(),
			//发放时间
			IssuedAt : time.Now().Unix(),
			//发放机构
			Issuer: "xutapd.dev",
			//主题
			Subject: "user Token",
		},
	}

	//选择HS的加密算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//解析Token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error){
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}

