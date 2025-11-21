package service

import (
	"book-manager/internal/mapper"

	"github.com/gin-gonic/gin"
)

type Service interface {
	//user
	UserCheck(rt *gin.Context)
	UserAdd(rt *gin.Context)
	//post
	Create(rt *gin.Context)
	GetList(rt *gin.Context)
	Update(rt *gin.Context)
	Delete(rt *gin.Context)
	//comment
	CreateComment(rt *gin.Context)
	GetListComment(rt *gin.Context)
}

type service struct {
	mappers mapper.Mapper
}

func ServiceInit(mapper mapper.Mapper) Service {
	return &service{mappers: mapper}
}
