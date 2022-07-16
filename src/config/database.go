package config

import (
	// "sample/domain/model"
	"github.com/katsukiniwa/practical-go-programming/src/domain/model"

	"github.com/jinzhu/gorm"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open("mysql", "user:password@tcp(sample_db)/sample?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(model.Task{})

	return db
}
