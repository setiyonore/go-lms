package repository

import (
	"go-lms/entities"

	"gorm.io/gorm"
)

type LateCharge interface {
	FindAll() ([]entities.LateCharge, error)
	FindById(id int) (entities.LateCharge, error)
	PayLateCharge(id int) error
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
func (r *latecharge) FindById(id int) (entities.LateCharge, error) {
	var data entities.LateCharge
	err := r.db.Where("late_charges.id", id).
		Select("id", "days_late", "is_pay", "id_borrowing").
		Preload("BookBorrowing", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "borrowing_date", "member_id")
		}).
		Preload("BookBorrowing.LibrarryMember", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).Find(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}
func (r *latecharge) PayLateCharge(id int) error {
	return nil
}
