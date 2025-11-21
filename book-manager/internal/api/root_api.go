package api

import (
	"book-manager/internal/mapper"
	"book-manager/internal/service"
	"book-manager/pkg/exception"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Apis(rt *gin.Engine, mappers mapper.Mapper) error {
	services := service.ServiceInit(mappers)
	//1.校验用户认证
	//魔板解析
	//rt.LoadHTMLFiles("web/template/login.tmpl")
	rt.LoadHTMLGlob("web/template/*")
	//1.注册页面，注册
	rt.Use(exception.ErrorHandler())
	//2.登录页面，登录
	login := rt.Group("/login")
	{
		login.GET("/", loginHtml)
		login.POST("/login", services.UserCheck)
	}
	//3.用户相关
	users := rt.Group("users")
	users.Use(JwtFilter())
	{
		users.POST("/add", services.UserAdd)
	}
	//4.post文章相关
	posts := rt.Group("/posts")
	posts.Use(JwtFilter())
	{
		//文章的创建功能
		posts.POST("/CreatePost", services.Create)
		//文章的读取功能
		posts.POST("/GetList", services.GetList)
		//文章的更新功能
		posts.POST("/update", services.Update)
		//文章的删除功能
		posts.POST("/delete", services.Delete)
	}
	//5.评论相关

	return nil
}

func loginHtml(rt *gin.Context) {
	rt.HTML(200, "login.tmpl", gin.H{})
}

// 认证中间件
func JwtFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		fmt.Println("init JwtFilter,auth=", auth)
		if auth == "" {
			c.JSON(200, gin.H{"msg": "jwt校验失败"})
			c.Abort() // 关键：必须调用Abort来终止后续处理
			return
		}
		token, _ := ParseToken(auth)
		c.Set("userId", (*token)["id"])
		c.Next()
	}
}

var secretKey = []byte("your_secret_key")

func ParseToken(tokenString string) (*jwt.MapClaims, error) {
	token, _ := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if maps, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		return maps, nil
	}
	return nil, nil
}
