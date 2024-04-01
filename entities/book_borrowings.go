package entities

import (
	"time"

	"gorm.io/gorm"
)

type BookBorrowings struct {
	ID                  uint                `json:"id"`
	BorrowingDate       string              `json:"borrowing_date"`
	ReturnDate          string              `json:"return_date"`
	UserID              int                 `json:"-"`
	MemberID            int                 `json:"-"`
	IsLateReturn        int                 `json:"is_late_return"`
	IsReturn            int                 `json:"is_return"`
	CreatedAt           time.Time           `json:"-"`
	UpdatedAt           time.Time           `json:"-"`
	DeletedAt           gorm.DeletedAt      `json:"-"`
	User                user                `gorm:"foreignkey:UserID"`
	LibrarryMember      LibrarryMember      `gorm:"foreignKey:MemberID"`
	BookBorrowingDetail []BookBorrowDetails `gorm:"foreignkey:IdBookBorrow"`
}
type BookBorrowingInput struct {
	BorrowingDate string     `json:"borrowing_date" validate:"required"`
	ReturnDate    string     `json:"return_date" validate:"required"`
	UserID        int        `json:"user_id" validate:"required"`
	MemberID      int        `json:"member_id" validate:"required"`
	Books         []BookItem `json:"books" validate:"required"`
}

type BookItem struct {
	IDBook int `json:"book_id" vaidate:"required"`
}

func FormatBookBorrowing(bookborrowing BookBorrowings) BookBorrowings {
	bookborrowingFormatter := BookBorrowings{}
	bookborrowingFormatter.ID = bookborrowing.ID
	bookborrowingFormatter.BorrowingDate = bookborrowing.BorrowingDate
	bookborrowingFormatter.ReturnDate = bookborrowing.ReturnDate
	bookborrowingFormatter.UserID = bookborrowing.UserID
	bookborrowingFormatter.MemberID = bookborrowing.MemberID
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
	ID   int    `json:"-"`
	Name string `json:"name"`
}

type LibrarryMember struct {
	ID   int    `json:"-"`
	Name string `json:"name"`
}
type BookBorrowDetails struct {
	ID           uint `json:"id"`
	IdBookBorrow uint `json:"-"`
	IdBook       uint `json:"-"`
	Book         book `gorm:"foreignkey:IdBook"`
}

type book struct {
	ID   uint   `json:"-"`
	Name string `json:"name"`
	Isbn string `json:"isbn"`
}
