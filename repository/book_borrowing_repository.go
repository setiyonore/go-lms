package repository

import (
	"go-lms/entities"

	"gorm.io/gorm"
)

type BookBorrowing interface {
	GetAll() ([]entities.BookBorrowings, error)
}

type bookborrowing struct {
	db *gorm.DB
}

func NewBookBorrowing(db *gorm.DB) *bookborrowing {
	return &bookborrowing{db: db}
}

func (r *bookborrowing) GetAll() ([]entities.BookBorrowings, error) {
	var bookBorrowings []entities.BookBorrowings
	err := r.db.Preload("User").Find(&bookBorrowings).Error
	if err != nil {
		return bookBorrowings, err
	}
	return bookBorrowings, nil
}
