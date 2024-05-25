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
	UpdateBorrowing(bookBorrowing entities.BookBorrowings) error
	DeleteBorrowingDetails(idBorrowing int) error
	BookReturn(bookBorrowing entities.BookBorrowings) error
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
			"user_id", "member_id").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Preload("LibrarryMember", func(db *gorm.DB) *gorm.DB {
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
		Select("id", "borrowing_date", "return_date", "is_late_return",
			"is_return", "user_id", "member_id").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Preload("LibrarryMember", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Preload("BookBorrowingDetail.ItemBook", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "isbn", "status", "id_book")
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

func (r *bookborrowing) UpdateBorrowing(bookBorrowing entities.BookBorrowings) error {
	err := r.db.Updates(&bookBorrowing).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *bookborrowing) DeleteBorrowingDetails(idBorrowing int) error {
	var data entities.BookBorrowDetails
	err := r.db.Where("id_book_borrow = ?", idBorrowing).Delete(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *bookborrowing) BookReturn(data entities.BookBorrowings) error {
	//TODO
	err := r.db.Model(&entities.BookBorrowings{}).Where("id = ?", data.ID).
		Updates(entities.BookBorrowings{
			IsLateReturn: data.IsLateReturn,
			IsReturn:     1,
		}).Error
	if err != nil {
		return err
	}
	return nil
}
