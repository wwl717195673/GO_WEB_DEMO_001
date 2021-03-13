package controller

import (
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory model.Category
	ctx.Bind(&requestCategory)

	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "数据名称验证错误，分类名称必填")
	}
	c.DB.Create(&requestCategory)
	response.Success(ctx, gin.H{"category": requestCategory}, "创建成功")
}

func (c CategoryController) Update(ctx *gin.Context) {
	panic("implement me")
}

func (c CategoryController) Show(ctx *gin.Context) {
	panic("implement me")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	panic("implement me")
}
func NewCategoryController() ICategoryController {
	db := common.GetDB()
	db.AutoMigrate(model.Category{})
	return CategoryController{DB: db}
}
