package interfaces

import "github.com/katsukiniwa/practical-go-programming/src/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	FindByName(string) ([]domain.User, error)
	FindAll() ([]domain.User, error)
}
