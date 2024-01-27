package entities

import (
	"time"

	"gorm.io/gorm"
)

type BookBorrowDetailsw struct {
	ID           uint           `json:"id"`
	IdBookBorrow uint           `json:"id_book_borrow"`
	IdBook       uint           `json:"id_book"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `json:"-"`
}
