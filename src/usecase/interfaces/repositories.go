package interfaces

import (
	"github.com/katsukiniwa/practical-go-programming/src/domain/model"
)

type UserRepository interface {
	Store(model.User) (int, error)
	FindByName(string) ([]model.User, error)
	FindAll() ([]model.User, error)
}
