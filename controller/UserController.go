package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/levnzzz/ginEssential/common"
	"github.com/levnzzz/ginEssential/dto"
	"github.com/levnzzz/ginEssential/model"
	"github.com/levnzzz/ginEssential/response"
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
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号不能少于11位")
		return
	}

	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于六位")
		return
	}

	if len(name) == 0 {
		name = "levnzzz"
	}

	log.Println(name, telephone, password)
	//判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}

	//创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		return
	}
	newUser := model.User{
		Name: name,
		Telephone: telephone,
		Password: string(hasedPassword),
	}

	DB.Create(&newUser)
	//返回结果
	response.Success(ctx, nil, "注册成功")
}

func Login(ctx *gin.Context) {
	db := common.DB
	//获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//数据验证
	if len(telephone) < 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号不能少于11位")
		return
	}

	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于六位")
		return
	}

	//判断手机号是否存在
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Fail(ctx, nil, "密码错误")
		return
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Println("token generate error : %v", err)
		return
	}
	//返回结果
	response.Success(ctx, gin.H{"token": token}, "登陆成功")
}

func Info(ctx *gin.Context)  {
	user, _ := ctx.Get("user")

	response.Success(ctx, gin.H{"user": dto.ToUserDto(user.(model.User))}, "登陆成功")
}



func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}


