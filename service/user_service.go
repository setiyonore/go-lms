package service

import (
	"go-lms/entities"
	"go-lms/repository"
)

type User interface {
	GetAllUser() ([]entities.User, error)
}

type user struct {
	userRepository repository.User
}

func NewUser(userRepository repository.User) *user {
	return &user{userRepository: userRepository}
}

func (u *user) GetAllUser() ([]entities.User, error) {
	users, err := u.userRepository.FindAll()
	if err != nil {
		return users, err
	}
	return users, err
}
