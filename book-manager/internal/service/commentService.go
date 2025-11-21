package service

import (
	"book-manager/internal/model"

	"github.com/gin-gonic/gin"
)

func (c *service) CreateComment(rt *gin.Context) {
	var com model.Comment
	rt.ShouldBind(&com)
	//2.数据入库
	if com.UserId > 0 && com.PostId > 0 {
		c.mappers.CreateComment(com)
		rt.JSON(200, gin.H{"msg": "success"})
	} else {

		rt.JSON(200, gin.H{"msg": "fail"})
	}
}

func (c *service) GetListComment(rt *gin.Context) {
	var com model.Comment
	rt.ShouldBind(&com)
	postId := com.PostId
	rt.JSON(200, c.mappers.GetByPostId(postId))
}
