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

func (b *book) FindAll() ([]entities.Book, error) {
	var books []entities.Book
	err := b.db.Preload("Publisher").Preload("Author").Find(&books).Error
	if err != nil {
		return books, err
	}
	return books, nil
}

func (b *book) FindById(id int) (entities.Book, error) {
	var book entities.Book
	err := b.db.Where("id", id).Preload("Publisher").Preload("Author").Find(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func (b *book) Save(book entities.Book) error {
	err := b.db.Save(&book).Error
	if err != nil {
		return err
	}
	return nil
}

func (b *book) Update(book entities.Book) error {
	err := b.db.Save(&book).Error
	if err != nil {
		return err
	}
	return nil
}

func (b *book) Delete(id int) error {
	var book entities.Book
	err := b.db.Where("id", id).First(&book).Error
	if err != nil {
		return err
	}
	err = b.db.Delete(&book).Error
	if err != nil {
		return err
	}
	return nil
}
