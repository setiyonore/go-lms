package repository

import (
	"go-lms/entities"
	"gorm.io/gorm"
)

type Author interface {
	FindAll() ([]entities.Author, error)
}
type author struct {
	db *gorm.DB
}

func NewAuthor(db *gorm.DB) *author {
	return &author{db}
}

func (a *author) FindAll() ([]entities.Author, error) {
	var authors []entities.Author
	err := a.db.Find(&authors).Error
	if err != nil {
		return authors, err
	}
	return authors, nil
}
