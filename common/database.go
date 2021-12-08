package common

import (
	"XUTAPD/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	username := "root"
	password := "7WzhXRTJdSEZWknE"
	host := "120.53.228.79"
	port := "3306"
	database := "xutapd"
	charset := "utf8mb4"
	loc := "Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))
	//需要全局db
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("connect Mysql Error: %v", err)
	}
	DB.AutoMigrate(&models.User{}) //自动迁移
	//传递全局变量
	db = DB
	return DB
}
//传递指针
func GetDB() *gorm.DB {
	return db
}