package model

import (
	"errors"
)

type User struct {
	ID           int
	FirstName    string
	LastName     string
	EmailAddress string
}

func NewUser(first_name, last_name string) (*User, error) {
	if first_name == "" {
		return nil, errors.New("first nameを入力してください")
	}

	user := &User{
		FirstName: first_name,
		LastName:  last_name,
	}

	return user, nil
}

func (t *User) SetUser(first_name, last_name string) error {
	if first_name == "" {
		return errors.New("first nameを入力してください")
	}

	t.FirstName = first_name
	t.LastName = last_name

	return nil
}
