package service

import (
	"go-lms/entities"
	"go-lms/repository"
)

type BookBorrowings interface {
	GetBookBorrowing() ([]entities.BookBorrowings, error)
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
