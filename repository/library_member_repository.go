package repository

import (
	"go-lms/entities"

	"gorm.io/gorm"
)

type LibraryMember interface {
	FindAll() ([]entities.LibrarryMembers, error)
	FindById(id int) (entities.LibrarryMembers, error)
}

type librarymember struct {
	db *gorm.DB
}

func NewLibraryMember(db *gorm.DB) *librarymember {
	return &librarymember{db: db}
}
func (r *librarymember) FindAll() ([]entities.LibrarryMembers, error) {
	var libraryMembers []entities.LibrarryMembers
	err := r.db.Find(&libraryMembers).Error
	if err != nil {
		return libraryMembers, err
	}
	return libraryMembers, nil
}

func (r *librarymember) FindById(id int) (entities.LibrarryMembers, error) {
	var libraryMember entities.LibrarryMembers
	err := r.db.Where("id", id).Find(&libraryMember).Error
	if err != nil {
		return libraryMember, err
	}
	return libraryMember, nil
}
