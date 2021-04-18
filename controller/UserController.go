package controller

import (
	"ginvue/model"
	"ginvue/utils"
	"ginvue/common"
	"log"
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
)

func isTelephoneExist(db *gorm.DB, tele string) bool {
	var user model.User
	db.Where("telephone = ?", tele).First(&user)
	
	return user.ID != 0
	// if user.ID != 0 {
	// 	return true
	// }
	// return false
}

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号必须为11位",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能少于6位",
		})
		return
	}
	// 如果名称没有传，则自动生成随机字符串
	if len(name) == 0 {
		name = utils.RandomString(10)
	}

	log.Println(name, telephone, password)

	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户已存在",
		})
		return
	}
	// 创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&newUser)
	// 返回结果
	ctx.JSON(200, gin.H{
		"msg": "注册成功",
	})
}
