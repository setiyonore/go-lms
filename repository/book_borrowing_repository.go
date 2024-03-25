package repository

import (
	"go-lms/entities"

	"gorm.io/gorm"
)

type BookBorrowing interface {
	GetAll() ([]entities.BookBorrowings, error)
	GetDetail(id int) (entities.BookBorrowings, error)
	SaveBorrowing(entities.BookBorrowings) (entities.BookBorrowings, error)
	SaveBorrowingDetails(idBorrowing int, dataDetails []entities.BookBorrowDetails) error
}

type bookborrowing struct {
	db *gorm.DB
}

func NewBookBorrowing(db *gorm.DB) *bookborrowing {
	return &bookborrowing{db: db}
}

func (r *bookborrowing) GetAll() ([]entities.BookBorrowings, error) {
	var bookBorrowings []entities.BookBorrowings
	err := r.db.
		Select("id", "borrowing_date", "return_date", "is_late_return", "is_return",
			"user_id").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Find(&bookBorrowings).Error
	if err != nil {
		return bookBorrowings, err
	}
	return bookBorrowings, nil
}

func (r *bookborrowing) GetDetail(id int) (entities.BookBorrowings, error) {
	var bookborrowing entities.BookBorrowings
	err := r.db.Where("id", id).
		Select("id", "borrowing_date", "return_date", "is_late_return", "is_return",
			"user_id").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Preload("BookBorrowingDetail.Book", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "isbn")
		}).
		Find(&bookborrowing).Error
	if err != nil {
		return bookborrowing, err
	}
	return bookborrowing, nil
}

func (r *bookborrowing) SaveBorrowing(bookborrowing entities.BookBorrowings) (entities.BookBorrowings, error) {
	err := r.db.Create(&bookborrowing).Error
	if err != nil {
		return bookborrowing, err
	}
	return bookborrowing, nil
}

func (r *bookborrowing) SaveBorrowingDetails(idBorrowing int, dataDetails []entities.BookBorrowDetails) error {
	for _, detail := range dataDetails {
		detail.IdBookBorrow = uint(idBorrowing)
		err := r.db.Create(&detail).Error
		if err != nil {
			return err
		}
	}
	return nil
}
