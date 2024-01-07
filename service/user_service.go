package service

import (
	"go-lms/entities"
	"go-lms/repository"
)

type User interface {
	GetAllUser() ([]entities.User, error)
	GetUserById(Id int) (entities.User, error)
	GetUserByEmail(Email string) (entities.User, error)
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

func (u *user) GetUserById(Id int) (entities.User, error) {
	user, err := u.userRepository.FindById(Id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *user) GetUserByEmail(Email string) (entities.User, error) {
	user, err := u.userRepository.FindByEmail(Email)
	if err != nil {
		return user, err
	}
	return user, nil
}