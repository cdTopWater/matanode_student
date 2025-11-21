package mapper

import (
	"book-manager/internal/model"
)

func (c *mapper) CreateComment(com model.Comment) int {
	c.db.Create(&com)
	return 1
}

func (c *mapper) GetByPostId(postId int) []model.Comment {
	var result []model.Comment
	c.db.Where("post_id = ? ", postId).Find(&result)
	return result
}
