package repository

import (
	"go-lms/entities"
	"gorm.io/gorm"
)

type Author interface {
	FindAll() ([]entities.Author, error)
	FindByID(ID int) (entities.Author, error)
}
type author struct {
	db *gorm.DB
}

func (a *author) FindByID(ID int) (entities.Author, error) {
	var author entities.Author
	err := a.db.Where("id = ?", ID).Find(&author).Error
	if err != nil {
		return author, err
	}
	return author, nil
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
