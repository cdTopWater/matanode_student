package model

type User struct {
	Id       int
	Username string `form:"userName"`
	Password string `form:"passWord"`
	Email    string `form:"email"`
}
