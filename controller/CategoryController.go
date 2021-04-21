package controller

import (
	"ginvue/confirm"
	"ginvue/model"
	"ginvue/repository"
	"ginvue/response"

	"strconv"

	"github.com/gin-gonic/gin"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	Repository repository.CategoryRespository
}

func NewCategoryController() ICategoryController {
	repository := repository.NewCategoryResp()
	repository.DB.AutoMigrate(model.Category{})
	return CategoryController{Repository: repository}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory confirm.CreateCategoryReq
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "数据验证错误，名称必须")
		return
	}
	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		panic(err)
	}

	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Update(ctx *gin.Context) {
	// 绑定body中的参数
	var requestCategory confirm.CreateCategoryReq
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "数据验证错误，名称必须")
		return
	}
	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	updateCategory, err := c.Repository.SelectById(categoryId)

	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}
	// 更新分类, 可以传三种参数 map / struct / name balue
	category, err := c.Repository.Update(updateCategory, requestCategory.Name)
	if err != nil {
		panic(err)
	}

	response.Success(ctx, gin.H{"category": category}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if err := c.Repository.DeleteById(categoryId); err != nil {
		response.Fail(ctx, nil, "删除失败")
		return
	}

	response.Success(ctx, nil, "删除成功")
}
