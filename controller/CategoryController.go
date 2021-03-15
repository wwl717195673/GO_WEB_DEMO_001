package controller

import (
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
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
		return
	}
	c.DB.Create(&requestCategory)
	response.Success(ctx, gin.H{"category": requestCategory}, "创建成功")
}

func (c CategoryController) Update(ctx *gin.Context) {
	var requestCategory model.Category
	ctx.Bind(&requestCategory)
	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "数据名称验证错误，分类名称必填")
		return
	}
	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var updateCategory model.Category
	if c.DB.First(&updateCategory, categoryId).RecordNotFound() {
		response.Fail(ctx, nil, "未找到该分组，请确保其是否创建")
		return
	}
	//更新分类
	//map
	//struct
	//name value
	c.DB.Model(&updateCategory).Update("name", requestCategory.Name)
	response.Success(ctx, gin.H{"category": updateCategory}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var category model.Category
	if c.DB.First(&category, categoryId).RecordNotFound() {
		response.Fail(ctx, nil, "未找到该分组，请确保其是否创建")
		return
	}
	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var category model.Category
	if c.DB.First(&category, categoryId).RecordNotFound() {
		response.Fail(ctx, nil, "未找到该分组，请确保其是否创建")
		return
	}
	if err := c.DB.Delete(model.Category{}, categoryId).Error; err != nil {
		response.Fail(ctx, nil, "删除失败")
		return
	}
	response.Success(ctx, nil, "删除成功")
}

func NewCategoryController() ICategoryController {
	db := common.GetDB()
	db.AutoMigrate(model.Category{})
	return CategoryController{DB: db}
}
