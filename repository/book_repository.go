package repository

import (
	"go-lms/entities"

	"gorm.io/gorm"
)

type Book interface {
	FindAll() ([]entities.Book, error)
	FindById(id int) (entities.Book, error)
	Save(book entities.Book) error
	Update(book entities.Book) error
	Delete(id int) error
}

type book struct {
	db *gorm.DB
}

func NewBook(db *gorm.DB) *book {
	return &book{db: db}
}

func (r *book) FindAll() ([]entities.Book, error) {
	var books []entities.Book
	err := r.db.Preload("Publisher").Preload("Author").Find(&books).Error
	if err != nil {
		return books, err
	}
	return books, nil
}

func (r *book) FindById(id int) (entities.Book, error) {
	var book entities.Book
	err := r.db.Where("id", id).Preload("Publisher").Preload("Author").Find(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func (r *book) Save(book entities.Book) error {
	err := r.db.Save(&book).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *book) Update(book entities.Book) error {
	err := r.db.Save(&book).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *book) Delete(id int) error {
	var book entities.Book
	err := r.db.Where("id", id).First(&book).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&book).Error
	if err != nil {
		return err
	}
	return nil
}
