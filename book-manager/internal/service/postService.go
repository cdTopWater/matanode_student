package service

import (
	"book-manager/internal/model"
	"book-manager/pkg/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (p *service) Create(rt *gin.Context) {
	userid, _ := rt.Get("userId")
	fmt.Println("userid=", userid)
	var post model.Post
	rt.ShouldBind(&post)
	c, _ := utils.AnyToInt64(userid)
	post.UserId = int(c)
	fmt.Println("post=", post)
	//数据入库
	if post.Title == "" || post.Content == "" {
		rt.JSON(200, gin.H{"msg": "fail"})
	} else {
		p.mappers.Create(post)
		rt.JSON(200, gin.H{"msg": "success"})
	}
}

func (p *service) GetList(rt *gin.Context) {
	var post model.Post
	rt.ShouldBind(&post)
	if post.Id > 0 {
		rt.JSON(200, []model.Post{p.mappers.GetbyId(post)})
	} else if post.Title != "" {
		rt.JSON(200, p.mappers.GetListByTitle(post.Title))
	}
}

func (p *service) Update(rt *gin.Context) {
	userid, _ := rt.Get("userId")
	fmt.Println("userid=", userid)
	var post model.Post
	rt.ShouldBind(&post)
	queryPost := p.mappers.GetbyId(post)
	if queryPost.UserId != userid {
		rt.JSON(200, gin.H{"msg": "只能更新自己的作品"})
	} else {
		p.mappers.Update(post)
		rt.JSON(200, gin.H{"msg": "更新成功"})
	}

}

func (p *service) Delete(rt *gin.Context) {
	userid, _ := rt.Get("userId")
	fmt.Println("userid=", userid)
	var post model.Post
	rt.ShouldBind(&post)
	queryPost := p.mappers.GetbyId(post)
	if queryPost.UserId != userid {
		rt.JSON(200, gin.H{"msg": "只能删除自己的作品"})
	} else {
		p.mappers.Delete(post.Id)
		rt.JSON(200, gin.H{"msg": "删除成功"})
	}

}
