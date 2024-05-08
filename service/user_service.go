package service

import (
	"errors"
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
	UpdateUser(inputID int, inputData entities.EditUserInput) error
	DeleteUser(id int) error
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
	// user.Role = Input.Role
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

func (u *user) UpdateUser(inputId int, inputData entities.EditUserInput) error {
	user, err := u.userRepository.FindById(inputId)
	if err != nil {
		return err
	}
	if user.ID == 0 {
		return errors.New("data not found")
	}
	user.Name = inputData.Name
	user.Email = inputData.Email
	// user.Role = inputData.Role
	if len(inputData.Password) > 0 {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(inputData.Password), bcrypt.MinCost)
		if err != nil {
			return err
		}
		user.Password = string(passwordHash)
	}
	err = u.userRepository.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *user) DeleteUser(id int) error {
	user, err := u.userRepository.FindById(id)
	if err != nil {
		return err
	}
	if user.ID == 0 {
		return errors.New("data not found")
	}
	err = u.userRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
