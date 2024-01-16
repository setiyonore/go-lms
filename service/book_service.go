package service

import (
	"errors"
	"go-lms/entities"
	"go-lms/repository"
)

type Book interface {
	GetBook() ([]entities.Book, error)
	GetBookById(id int) (entities.Book, error)
}

type book struct {
	bookRepository repository.Book
}

func NewBook(bookRepository repository.Book) *book {
	return &book{bookRepository: bookRepository}
}

func (b *book) GetBook() ([]entities.Book, error) {
	books, err := b.bookRepository.FindAll()
	if err != nil {
		return books, err
	}
	return books, nil
}

func (b *book) GetBookById(id int) (entities.Book, error) {
	book, err := b.bookRepository.FindById(id)
	if err != nil {
		return book, err
	}
	if book.ID == 0 {
		err = errors.New("data not found")
		return book, err
	}
	return book, nil
}
