package exception

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrDatabaseConnection = Error(http.StatusInternalServerError, "数据库连接失败", "")
	ErrUserNotFound       = Error(http.StatusNotFound, "用户不存在", "")
	ErrInvalidCredentials = Error(http.StatusUnauthorized, "用户名或密码错误", "")
	ErrArticleNotFound    = Error(http.StatusNotFound, "文章不存在", "")
	ErrCommentNotFound    = Error(http.StatusNotFound, "评论不存在", "")
	ErrUnauthorized       = Error(http.StatusUnauthorized, "未授权访问", "")
	ErrBadRequest         = Error(http.StatusBadRequest, "请求参数错误", "")
)

type GlobleError struct {
	Code    int
	Message string
	Details string
}

func (e *GlobleError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s, Details: %s", e.Code, e.Message, e.Details)
}

func Error(code int, massage string, details string) *GlobleError {
	return &GlobleError{
		Code:    code,
		Message: massage,
		Details: details,
	}
}

// 处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			er := ctx.Errors.Last().Err
			if errs, ok := er.(*GlobleError); ok {
				ctx.JSON(errs.Code, gin.H{
					"msg": errs.Message,
				})
			} else {
				log.Printf("未知错误: %v", er)
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": "内部服务器错误",
				})
			}
		}
	}
}
