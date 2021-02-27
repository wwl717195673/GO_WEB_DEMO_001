package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"math/rand"
	"net/http"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(11);not null;unique"`
	Password  string `gorm:size:255;not null`
}

func main() {
	db := initDB()
	defer db.Close()
	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {
		fmt.Println(c.Keys)
		//获取参数
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")
		//数据验证
		if len(telephone) != 11 {
			fmt.Println(telephone)
			fmt.Println(len(telephone))
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "电话号码不符合格式"})
			return
		}
		if len(password) < 6 {
			fmt.Println(len(password))
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码必须6位以上"})
			return
		}
		if len(name) == 0 {
			name = randomName(10)
		}

		log.Println(name, telephone, password)

		//查询手机号是否存在
		if isTelephoneExist(db, telephone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号已经存在"})
			return
		}

		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "register success"})
	})

	panic(r.Run())
}

//generate user's name if he doesn't have one
func randomName(n int) string {
	var letters = []byte("asdfghjklzxcvbnmQWERTYUIOPZXCVBNM")
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

//initial database
func initDB() *gorm.DB {
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

	db.AutoMigrate(&User{})

	return db
}

func isTelephoneExist(db *gorm.DB, tel string) bool {
	var user User
	db.Where("telephone = ?", tel).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
