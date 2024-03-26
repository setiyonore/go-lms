package service

import (
	"errors"
	"go-lms/entities"
	"go-lms/repository"
)

type BookBorrowings interface {
	GetBookBorrowing() ([]entities.BookBorrowings, error)
	GetDetailBorrowing(id int) (entities.BookBorrowings, error)
	AddBookBorrowing(input entities.BookBorrowingInput) (string, error)
}

type bookborrowings struct {
	bookBorrowingsRepository repository.BookBorrowing
	bookRepository           repository.Book
}

func NewBookBorrowing(bookBorrowingRepository repository.BookBorrowing, bookRepository repository.Book) *bookborrowings {
	return &bookborrowings{bookBorrowingsRepository: bookBorrowingRepository, bookRepository: bookRepository}
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

func (s *bookborrowings) AddBookBorrowing(input entities.BookBorrowingInput) (string, error) {
	var message string
	details := make([]entities.BookBorrowDetails, len(input.Books))
	for i, book := range input.Books {
		details[i].IdBook = uint(book.IDBook)
		count := s.bookRepository.CheckBookAvalable(book.IDBook)
		if count > 0 {
			book, _ := s.bookRepository.FindById(book.IDBook)
			message = "book: " + book.Name + ", isbn: " + book.Isbn + " not available"
			return message, errors.New("book not available")
		}
	}
	BookBorowwing := entities.BookBorrowings{}
	BookBorowwing.BorrowingDate = input.BorrowingDate
	BookBorowwing.ReturnDate = input.ReturnDate
	BookBorowwing.UserID = input.UserID
	BookBorowwing.MemberID = input.MemberID
	result, err := s.bookBorrowingsRepository.SaveBorrowing(BookBorowwing)
	if err != nil {
		message = "failed save book borrowing"
		return message, err
	}
	// fmt.Println("form service", details)
	err = s.bookBorrowingsRepository.SaveBorrowingDetails(int(result.ID), details)
	if err != nil {
		message = "failed save book borrowing"
		return message, err
	}
	message = "success save book borrowing"
	return message, nil
}
