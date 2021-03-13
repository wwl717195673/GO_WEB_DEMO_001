package common

import (
	"fmt"
	"ginEssential/model"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

//initial database
func InitDB() *gorm.DB {
	drivename := viper.GetString("datasource.driveName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
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
