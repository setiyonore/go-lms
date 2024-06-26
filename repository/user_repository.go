package repository

import (
	"go-lms/entities"

	"gorm.io/gorm"
)

type User interface {
	FindAll() ([]entities.User, error)
	FindById(Id int) (entities.User, error)
	FindByEmail(Email string) (entities.User, error)
	Save(User entities.User) error
	Update(User entities.User) error
	Delete(id int) error
}

type user struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *user {
	return &user{db: db}
}

func (u *user) FindAll() ([]entities.User, error) {
	var users []entities.User
	err := u.db.
		Select("id", "name", "email").
		Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}
func (u *user) FindById(Id int) (entities.User, error) {
	var user entities.User
	err := u.db.Where("id", Id).
		Select("id", "name", "email").
		Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *user) FindByEmail(Email string) (entities.User, error) {
	var user entities.User
	err := u.db.Where("email", Email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func (u *user) Save(User entities.User) error {
	err := u.db.Save(&User).Error
	if err != nil {
		return err
	}
	return nil
}
func (u *user) Update(User entities.User) error {
	err := u.db.Updates(&User).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *user) Delete(id int) error {
	var user entities.User
	err := u.db.Where("id", id).First(&user).Error
	if err != nil {
		return err
	}
	err = u.db.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil

}
