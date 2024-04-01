package entities

import (
	"time"

	"gorm.io/gorm"
)

type LateCharge struct {
	ID            uint           `json:"id" gorm:"primary_key"`
	IdBorrowing   uint           `json:"id_borrowing"`
	DaysLate      int            `json:"days_late"`
	IsPay         int            `json:"is_pay"`
	CreatedAt     time.Time      `json:"-"`
	UpdateAt      time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `json:"-"`
	BookBorrowing bookBorrowing  `gorm:"foreignKey:IdBorrowing"`
}

type bookBorrowing struct {
	ID             uint           `json:"-"`
	BorrowingDate  string         `json:"borrowing_date"`
	MemberID       int            `json:"-"`
	LibrarryMember LibrarryMember `gorm:"foreignKey:MemberID"`
}

func FormatLateCharge(latecharge LateCharge) LateCharge {
	lateChargeFormatter := LateCharge{}
	lateChargeFormatter.ID = latecharge.ID
	lateChargeFormatter.DaysLate = latecharge.DaysLate
	lateChargeFormatter.IdBorrowing = latecharge.IdBorrowing
	lateChargeFormatter.IsPay = latecharge.IsPay
	lateChargeFormatter.BookBorrowing.BorrowingDate = latecharge.BookBorrowing.BorrowingDate
	lateChargeFormatter.BookBorrowing.LibrarryMember.Name = latecharge.BookBorrowing.LibrarryMember.Name
	return lateChargeFormatter
}

func FormatLateCharges(latecharges []LateCharge) []LateCharge {
	latechargesFormatter := []LateCharge{}
	for _, latecharge := range latecharges {
		latechargeFormatter := FormatLateCharge(latecharge)
		latechargesFormatter = append(latechargesFormatter, latechargeFormatter)
	}
	return latechargesFormatter
}
