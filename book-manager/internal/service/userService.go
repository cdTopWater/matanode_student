package service

import (
	"book-manager/internal/model"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (service *service) UserCheck(rt *gin.Context) {
	userName := rt.PostForm("userName")
	fmt.Printf("userName=", userName)
	var loginUser = model.User{}
	rt.ShouldBind(&loginUser)
	fmt.Printf("loginUser=", loginUser)

	var user model.User = service.mappers.GetUserByIdOrName(model.User{Username: loginUser.Username})
	if user.Id <= 0 {
		// 没查询出来，返回注册页
		rt.HTML(200, "regist.tmpl", gin.H{})
	} else {
		//加密后比较
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(loginUser.Password), bcrypt.DefaultCost)
		if err != nil {
			//跳转注册
			rt.HTML(200, "regist.tmpl", gin.H{})
		}
		if user.Password == string(hashedPassword) {
			//登录成功,并返回jwt
			rt.HTML(200, "login.tmpl", gin.H{"token": genToken(user), "msg": "登录成功"})
		} else {
			//账号或密码错误
			rt.HTML(200, "login.tmpl", gin.H{"msg": "账号或密码错误"})
		}
	}
	rt.JSON(200, gin.H{})
}

// 生成jwt
func genToken(user model.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.Id,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := token.SignedString([]byte("your_secret_key"))
	return tokenString
}

func (service *service) UserAdd(rt *gin.Context) {
	var user model.User
	rt.ShouldBind(&user)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	} else {
		user.Password = string(hashedPassword)
		service.mappers.UserAdd(user)
		return
	}
}
