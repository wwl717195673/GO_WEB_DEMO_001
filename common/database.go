package common

import (
	"fmt"
	"ginEssential/model"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

//initial database
func InitDB() *gorm.DB {
	drivename := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginEssential"
	username := "root"
	password := "123456"
	charset := "utf8mb4"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(drivename, args)
	if err != nil {
		panic("failed to connect database,reason:" + err.Error())
	}

	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
