package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/levnzzz/ginEssential/common"
	"github.com/levnzzz/ginEssential/model"
	"log"
	"net/http"
)


func Register(ctx *gin.Context) {
	//获取参数
	DB := common.DB
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code" : 422, "msg" : "手机号必须为11位"})
		return
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code" : 422, "msg" : "密码必须6位数"})
		return
	}

	if len(name) == 0 {
		name = "levnzzz"
	}

	log.Println(name, telephone, password)
	//判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code" : 422, "msg" : "用户已存在"})
		return
	}

	//创建用户
	newUser := model.User{
		Name: name,
		Telephone: telephone,
		Password: password,
	}

	DB.Create(&newUser)
	//返回结果
	ctx.JSON(200, gin.H{
		"msg": "注册成功",
	})
}


func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

