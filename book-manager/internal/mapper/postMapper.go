package mapper

import (
	"book-manager/internal/model"
)

func (p *mapper) Create(u model.Post) int {
	p.db.Create(&u)
	return 1
}

func (p *mapper) GetbyId(u model.Post) model.Post {
	p.db.Find(&u)
	return u
}

func (p *mapper) GetListByTitle(title string) (result []model.Post) {
	p.db.Where("title like %?%", title).Find(&result)
	return
}

func (p *mapper) Update(u model.Post) int {
	p.db.Model(&model.Post{}).Where("id =?", u.Id).Updates(&u)
	return 1
}

func (p *mapper) Delete(id int) int {
	p.db.Delete(id)
	return 1
}
