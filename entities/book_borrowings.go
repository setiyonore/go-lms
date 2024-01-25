package entities

import (
	"time"

	"gorm.io/gorm"
)

type BookBorrowings struct {
	ID            uint           `json:"id"`
	BorrowingDate string         `json:"borrowing_date"`
	ReturnDate    string         `json:"return_date"`
	UserID        int            `json:"-"`
	IsLateReturn  bool           `json:"is_late_return"`
	IsReturn      bool           `json:"is_return"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `json:"-"`
	User          user           `gorm:"foreignkey:UserID"`
}

func FormatBookBorrowing(bookborrowing BookBorrowings) BookBorrowings {
	bookborrowingFormatter := BookBorrowings{}
	bookborrowingFormatter.ID = bookborrowing.ID
	bookborrowingFormatter.BorrowingDate = bookborrowing.BorrowingDate
	bookborrowingFormatter.ReturnDate = bookborrowing.ReturnDate
	bookborrowingFormatter.UserID = bookborrowing.UserID
	bookborrowingFormatter.IsLateReturn = bookborrowing.IsLateReturn
	bookborrowingFormatter.IsReturn = bookborrowing.IsReturn
	return bookborrowing
}

func FormatBookBorrowings(bookborrowings []BookBorrowings) []BookBorrowings {
	bookborrowingsformatter := []BookBorrowings{}
	for _, bookborrowing := range bookborrowings {
		bookborrowingformatter := FormatBookBorrowing(bookborrowing)
		bookborrowingsformatter = append(bookborrowingsformatter, bookborrowingformatter)
	}
	return bookborrowingsformatter
}

type user struct {
	ID        int            `json:"-"`
	Name      string         `json:"name"`
	Email     string         `json:"-"`
	Password  string         `json:"-"`
	Role      int            `json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
