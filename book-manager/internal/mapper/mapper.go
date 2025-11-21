package mapper

import (
	"book-manager/internal/model"

	"gorm.io/gorm"
)

type Mapper interface {
	//user
	GetUserByIdOrName(u model.User) model.User
	GetUserByNameAndPsw(u model.User) model.User
	UserAdd(u model.User) int
	//post
	Create(u model.Post) int
	GetbyId(u model.Post) model.Post
	GetListByTitle(title string) []model.Post
	Update(u model.Post) int
	Delete(id int) int
	//comment
	CreateComment(com model.Comment) int
	GetByPostId(postId int) []model.Comment
}

type mapper struct {
	db *gorm.DB
}

func MapperInit(dbs *gorm.DB) Mapper {
	return &mapper{db: dbs}
}
