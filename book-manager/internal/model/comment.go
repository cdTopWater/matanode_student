package model

import "time"

type Comment struct {
	Id        int
	Content   string
	UserId    int `gorm:"index"`
	PostId    int `gorm:"index"`
	CreatedAt time.Time
}
