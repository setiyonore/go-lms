package repository

import (
	"go-lms/entities"

	"gorm.io/gorm"
)

type Book interface {
	FindAll() ([]entities.Book, error)
	FindById(id int) (entities.Book, error)
	Save(book entities.Book) error
	Update(book entities.Book) error
	Delete(id int) error
	CheckBookAvalable(id int) int64
	GetItemBook(id int) (entities.Book, error)
}

type book struct {
	db *gorm.DB
}

func NewBook(db *gorm.DB) *book {
	return &book{db: db}
}

func (r *book) FindAll() ([]entities.Book, error) {
	var books []entities.Book
	err := r.db.
		Select("id", "name", "publisher_id", "author_id", "isbn",
			"year_of_publication").
		Preload("Publisher", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Preload("Author", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Find(&books).Error
	if err != nil {
		return books, err
	}
	return books, nil
}

func (r *book) FindById(id int) (entities.Book, error) {
	var book entities.Book
	err := r.db.Where("id", id).
		Select("id", "name", "description", "publisher_id", "author_id", "isbn",
			"year_of_publication", "img_url_thumbnail", "img_url_cover", "is_available").
		Preload("Publisher", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Preload("Author", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Find(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func (r *book) Save(book entities.Book) error {
	err := r.db.Save(&book).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *book) Update(book entities.Book) error {
	err := r.db.Updates(&book).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *book) Delete(id int) error {
	var book entities.Book
	err := r.db.Where("id", id).First(&book).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&book).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *book) CheckBookAvalable(id int) int64 {
	// 1 = not available
	//0 = available
	var count int64
	err := r.db.Model(&entities.BookBorrowDetails{}).
		Where("id_book = ?", id).Count(&count).Error
	if err != nil {
		return 100 //auto not available
	}
	return count
}

func (r *book) GetItemBook(id int) (entities.Book, error) {
	var book entities.Book

	err := r.db.Where("id = ?", id).
		Preload("Publisher", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Preload("Author", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Preload("ItemBooks").
		Find(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}
