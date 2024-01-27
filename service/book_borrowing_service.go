package service

import (
	"go-lms/entities"
	"go-lms/repository"
)

type BookBorrowings interface {
	GetBookBorrowing() ([]entities.BookBorrowings, error)
	GetDetailBorrowing(id int) (entities.BookBorrowings, error)
}

type bookborrowings struct {
	bookBorrowingsRepository repository.BookBorrowing
}

func NewBookBorrowing(bookBorrowingRepository repository.BookBorrowing) *bookborrowings {
	return &bookborrowings{bookBorrowingsRepository: bookBorrowingRepository}
}

func (s *bookborrowings) GetBookBorrowing() ([]entities.BookBorrowings, error) {
	bookborrowings, err := s.bookBorrowingsRepository.GetAll()
	if err != nil {
		return bookborrowings, err
	}
	return bookborrowings, nil
}

func (s *bookborrowings) GetDetailBorrowing(id int) (entities.BookBorrowings, error) {
	bookborrowing, err := s.bookBorrowingsRepository.GetDetail(id)
	if err != nil {
		return bookborrowing, err
	}
	return bookborrowing, nil
}
