package controller

import (
	"fmt"
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	DB := common.GetDB()
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
		name = util.RandomName(10)
	}

	log.Println(name, telephone, password)

	//查询手机号是否存在
	if isTelephoneExist(DB, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号已经存在"})
		return
	}

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&newUser)
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "register success"})
}

func isTelephoneExist(db *gorm.DB, tel string) bool {
	var user model.User
	db.Where("telephone = ?", tel).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
