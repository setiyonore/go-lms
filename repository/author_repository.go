package repository

import (
	"go-lms/entities"

	"gorm.io/gorm"
)

type Author interface {
	FindAll() ([]entities.Author, error)
	FindByID(ID int) (entities.Author, error)
	Save(author entities.Author) error
	Update(author entities.Author) error
	Delete(ID int) error
}
type author struct {
	db *gorm.DB
}

func (a *author) FindByID(ID int) (entities.Author, error) {
	var author entities.Author
	err := a.db.Where("id = ?", ID).
		Select("id", "name").
		Find(&author).Error
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
	err := a.db.
		Select("id", "name").
		Find(&authors).Error
	if err != nil {
		return authors, err
	}
	return authors, nil
}

func (a *author) Save(author entities.Author) error {
	err := a.db.Save(&author).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *author) Update(author entities.Author) error {
	err := a.db.Updates(&author).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *author) Delete(ID int) error {
	var author entities.Author
	err := a.db.Where("Id = ?", ID).First(&author).Error
	if err != nil {
		return err
	}
	err = a.db.Delete(&author).Error
	if err != nil {
		return err
	}
	return nil
}
