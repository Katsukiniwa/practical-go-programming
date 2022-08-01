package gateway

import (
	"github.com/jinzhu/gorm"
	"github.com/katsukiniwa/practical-go-programming/src/domain/model"
)

type (
	UserRepository struct {
		Conn *gorm.DB
	}

	User struct {
		gorm.Model
		FirstName    string `gorm:"size:20;not null"`
		EmailAddress string `gorm:"size:100;not null"`
		Age          int    `gorm:"type:smallint"`
	}
)

func (r *UserRepository) Store(u model.User) (id int, err error) {
	user := &User{
		FirstName:    u.FirstName,
		EmailAddress: u.EmailAddress,
	}

	if err = r.Conn.Create(user).Error; err != nil {
		return
	}

	return int(user.ID), nil
}

func (r *UserRepository) FindByName(name string) (d []model.User, err error) {
	users := []User{}
	if err = r.Conn.Where("name = ?", name).Find(&users).Error; err != nil {
		return
	}

	n := len(users)
	d = make([]model.User, n)
	for i := 0; i < n; i++ {
		d[i].ID = int(users[i].ID)
		d[i].FirstName = users[i].FirstName
		d[i].EmailAddress = users[i].EmailAddress
	}
	return
}

func (r *UserRepository) FindAll() (d []model.User, err error) {
	users := []User{}
	if err = r.Conn.Find(&users).Error; err != nil {
		return
	}

	n := len(users)
	d = make([]model.User, n)
	for i := 0; i < n; i++ {
		d[i].ID = int(users[i].ID)
		d[i].FirstName = users[i].FirstName
		d[i].EmailAddress = users[i].EmailAddress
	}
	return
}
