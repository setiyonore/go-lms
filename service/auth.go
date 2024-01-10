package service

import (
	"errors"
	"go-lms/entities"
	"go-lms/repository"

	"golang.org/x/crypto/bcrypt"
)

type Auth interface {
	Login(input entities.LoginInput) (entities.User, error)
}
type auth struct {
	userRepository repository.User
}

func NewAuth(userRepository repository.User) *auth {
	return &auth{userRepository: userRepository}
}

func (u *auth) Login(input entities.LoginInput) (entities.User, error) {
	email := input.Email
	password := input.Password

	user, err := u.userRepository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}
