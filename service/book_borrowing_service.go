package service

import (
	"go-lms/entities"
	"go-lms/repository"
)

type BookBorrowings interface {
	GetBookBorrowing() ([]entities.BookBorrowings, error)
	GetDetailBorrowing(id int) (entities.BookBorrowings, error)
	AddBookBorrowing(input entities.BookBorrowingInput) error
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

func (s *bookborrowings) AddBookBorrowing(input entities.BookBorrowingInput) error {
	BookBorowwing := entities.BookBorrowings{}
	BookBorowwing.BorrowingDate = input.BorrowingDate
	BookBorowwing.ReturnDate = input.ReturnDate
	BookBorowwing.UserID = input.UserID
	BookBorowwing.MemberID = input.MemberID
	result, err := s.bookBorrowingsRepository.SaveBorrowing(BookBorowwing)
	if err != nil {
		return err
	}
	details := make([]entities.BookBorrowDetails, len(input.Books))
	for i, book := range input.Books {
		details[i].IdBook = uint(book.IDBook) // Potential issue here
	}
	// fmt.Println("form service", details)
	err = s.bookBorrowingsRepository.SaveBorrowingDetails(int(result.ID), details)
	if err != nil {
		return err
	}
	return nil
}
