package repository

import (
	"go-lms/entities"

	"gorm.io/gorm"
)

type ItemBook interface {
	FindAll() ([]entities.ItemBook, error)
	FIndById(id int) (entities.ItemBook, error)
	FIndByIdBook(id int) ([]entities.ItemBook, error)
	FIndByIdBookAvailable(id int) ([]entities.ItemBook, error)
	Save(itemBook entities.ItemBook) error
	Update(itemBook entities.ItemBook) error
	UpdateStatus(id int, status int) error
	Delete(id int) error
}

type itemBook struct {
	db *gorm.DB
}

func NewItemBook(db *gorm.DB) *itemBook {
	return &itemBook{db: db}
}

func (r *itemBook) FindAll() ([]entities.ItemBook, error) {
	var itemBook []entities.ItemBook
	err := r.db.Select("id", "isbn", "id_book", "status").Find(&itemBook).Error
	if err != nil {
		return itemBook, err
	}
	return itemBook, nil
}

func (r *itemBook) FIndById(id int) (entities.ItemBook, error) {
	var itemBook entities.ItemBook
	err := r.db.Where("id", id).
		Select("id", "isbn", "id_book", "status").Find(&itemBook).Error
	if err != nil {
		return itemBook, err
	}
	return itemBook, nil
}

func (r *itemBook) Save(itemBook entities.ItemBook) error {
	err := r.db.Save(&itemBook).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *itemBook) Update(itemBook entities.ItemBook) error {
	err := r.db.Updates(&itemBook).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *itemBook) UpdateStatus(id int, status int) error {
	itemBook := entities.ItemBook{}
	err := r.db.Model(&itemBook).Where("id", id).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *itemBook) Delete(id int) error {
	var itemBook entities.ItemBook

	err := r.db.Where("id", id).First(&itemBook).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&itemBook).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *itemBook) FIndByIdBook(id int) ([]entities.ItemBook, error) {
	var itemBooks []entities.ItemBook
	err := r.db.Where("id_book", id).Select("id", "isbn", "id_book", "status").
		Find(&itemBooks).Error
	if err != nil {
		return itemBooks, err
	}
	return itemBooks, nil
}

func (r *itemBook) FIndByIdBookAvailable(id int) ([]entities.ItemBook, error) {
	var itemBooks []entities.ItemBook
	err := r.db.Where("id_book", id).
		Where("status", 1).
		Select("id", "isbn", "id_book", "status").
		Find(&itemBooks).Error
	if err != nil {
		return itemBooks, err
	}
	return itemBooks, err
}
