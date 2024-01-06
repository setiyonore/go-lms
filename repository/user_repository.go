package repository

import (
	"go-lms/entities"
	"gorm.io/gorm"
)

type User interface {
	FindAll() ([]entities.User, error)
}

type user struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *user {
	return &user{db: db}
}

func (u *user) FindAll() ([]entities.User, error) {
	var users []entities.User
	err := u.db.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}
