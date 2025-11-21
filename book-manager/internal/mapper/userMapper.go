package mapper

import (
	"book-manager/internal/model"
)

func (mapper *mapper) GetUserByIdOrName(u model.User) model.User {
	if u.Id > 0 {
		mapper.db.Find(&u)
	} else if u.Username != "" {
		mapper.db.Where("user_name = ?", u.Username).Find(&u)
	}
	return u
}

func (mapper *mapper) GetUserByNameAndPsw(u model.User) model.User {
	mapper.db.Where("user_name = ? and pass_word = ?", u.Username, u.Password).Find(&u)
	return u
}

func (mapper *mapper) UserAdd(u model.User) int {
	mapper.db.Create(&u)
	return 1
}
