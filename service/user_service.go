package service

import (
	"go-lms/entities"
	"go-lms/repository"
	"golang.org/x/crypto/bcrypt"
)

type User interface {
	GetAllUser() ([]entities.User, error)
	GetUserById(Id int) (entities.User, error)
	GetUserByEmail(Email string) (entities.User, error)
	AddUser(Input entities.AddUserInput) error
	IsEmailAvailable(Email string) (bool, error)
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

func (u *user) AddUser(Input entities.AddUserInput) error {
	user := entities.User{}
	user.Name = Input.Name
	user.Email = Input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(Input.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	err = u.userRepository.Save(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *user) IsEmailAvailable(Email string) (bool, error) {
	user, err := u.userRepository.FindByEmail(Email)
	if err != nil {
		return false, err
	}
	if user.ID == 0 {
		return true, nil
	}
	return false, nil
}
