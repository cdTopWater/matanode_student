package model

import "time"

type Post struct {
	Id        int
	Title     string
	Content   string
	UserId    int `grom:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// type Post struct {
// 	Id         int
// 	Title      string
// 	Content    string
// 	UserId     int       `gorm:"not null;index"` // 外键
// 	Comments   []Comment `gorm:"foreignKey:PostID"`
// 	CommentNum int       `gorm:"default:0"`
// 	State      string
// }
