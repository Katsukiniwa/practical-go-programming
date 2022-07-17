package usecase

import (
	"github.com/katsukiniwa/practical-go-programming/src/domain"
	"github.com/katsukiniwa/practical-go-programming/src/usecase/interfaces"
)

type UserInteractor struct {
	UserRepository interfaces.UserRepository
	Logger         interfaces.Logger
}

func (i *UserInteractor) Add(u domain.User) (int, error) {
	i.Logger.Log("store user!")
	return i.UserRepository.Store(u)
}
