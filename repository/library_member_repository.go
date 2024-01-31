package repository

import (
	"go-lms/entities"

	"gorm.io/gorm"
)

type LibraryMember interface {
	FindAll() ([]entities.LibrarryMembers, error)
	FindById(id int) (entities.LibrarryMembers, error)
	FindByName(name string) (entities.LibrarryMembers, error)
	Save(librarryMember entities.LibrarryMembers) error
	Update(librarymember entities.LibrarryMembers) error
	Delete(id int) error
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

func (r *librarymember) FindByName(name string) (entities.LibrarryMembers, error) {
	var libraryMember entities.LibrarryMembers
	err := r.db.Where("name LIKE ?", "%"+name+"%").Find(&libraryMember).Error
	if err != nil {
		return libraryMember, err
	}
	return libraryMember, nil
}

func (r *librarymember) Save(librarryMember entities.LibrarryMembers) error {
	err := r.db.Save(&librarryMember).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *librarymember) Update(libraryMember entities.LibrarryMembers) error {
	err := r.db.Save(&libraryMember).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *librarymember) Delete(id int) error {
	var libraryMember entities.LibrarryMembers
	err := r.db.Where("id", id).First(&libraryMember).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&libraryMember).Error
	if err != nil {
		return err
	}
	return nil
}
