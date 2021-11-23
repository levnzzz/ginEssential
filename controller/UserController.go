package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/levnzzz/ginEssential/common"
	"github.com/levnzzz/ginEssential/model"
	"golang.org/x/crypto/bcrypt"
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
	if len(telephone) < 11 {
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
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密错误"})
	}
	newUser := model.User{
		Name: name,
		Telephone: telephone,
		Password: string(hasedPassword),
	}

	DB.Create(&newUser)
	//返回结果
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg": "注册成功",
	})
}

func Login(ctx *gin.Context) {
	db := common.DB
	//获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//数据验证
	if len(telephone) < 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code" : 422, "msg" : "手机号必须为11位"})
		return
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code" : 422, "msg" : "密码必须6位数"})
		return
	}

	//判断手机号是否存在
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code" : 422, "msg" : "用户不存在"})
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code" : 400, "msg": "密码错误"})
		return
	}
	//发放token
	token := "11"
	//返回结果
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{"token" : token},
		"msg": "登陆成功",
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

