package repository

import (
	"go-lms/entities"

	"gorm.io/gorm"
)

type LateCharge interface {
	FindAll() ([]entities.LateCharge, error)
}

type latecharge struct {
	db *gorm.DB
}

func NewLateCharge(db *gorm.DB) *latecharge {
	return &latecharge{db: db}
}

func (r *latecharge) FindAll() ([]entities.LateCharge, error) {
	var lateCharges []entities.LateCharge
	err := r.db.Where("late_charges.is_pay", 0).
		Select("id", "days_late", "is_pay", "id_borrowing").
		Preload("BookBorrowing", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "borrowing_date", "member_id")
		}).
		Preload("BookBorrowing.LibrarryMember", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).Find(&lateCharges).Error
	if err != nil {
		return nil, err
	}
	return lateCharges, nil
}
