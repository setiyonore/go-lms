package repository

import (
	"go-lms/entities"

	"gorm.io/gorm"
)

type BookBorrowing interface {
	GetAll() ([]entities.BookBorrowings, error)
	GetDetail(id int) (entities.BookBorrowings, error)
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

func (r *bookborrowing) GetDetail(id int) (entities.BookBorrowings, error) {
	var bookborrowing entities.BookBorrowings
	err := r.db.Where("id", id).Preload("User").Preload("BookBorrowingDetail.Book").Find(&bookborrowing).Error
	if err != nil {
		return bookborrowing, err
	}
	return bookborrowing, nil
}
