package service

import (
	"errors"
	"go-lms/entities"
	"go-lms/repository"
)

type Book interface {
	GetBook() ([]entities.Book, error)
	GetBookById(id int) (entities.Book, error)
	AddBook(input entities.AddBookInput) error
	UpdateBook(inputId int, inputData entities.AddBookInput) error
	DeleteBook(id int) error
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

func (b *book) AddBook(input entities.AddBookInput) error {
	book := entities.Book{}
	book.Name = input.Name
	book.Description = input.Description
	book.AuthorID = uint(input.AuthorID)
	book.PublisherID = uint(input.PublisherID)
	book.Isbn = input.Isbn
	book.YearOfPublication = input.YearOfPublication
	book.ImgUrlThumbnail = input.ImgUrlThumbnail
	book.ImgUrlCover = input.ImgUrlCover
	err := b.bookRepository.Save(book)
	if err != nil {
		return err
	}
	return nil
}

func (b *book) UpdateBook(inputId int, inputData entities.AddBookInput) error {
	book, err := b.bookRepository.FindById(inputId)
	if err != nil {
		return err
	}
	if book.ID == 0 {
		err = errors.New("data not found")
		return err
	}
	book = entities.Book{}
	book.ID = uint(inputId)
	book.Name = inputData.Name
	book.Description = inputData.Description
	book.AuthorID = uint(inputData.AuthorID)
	book.PublisherID = uint(inputData.PublisherID)
	book.Isbn = inputData.Isbn
	book.YearOfPublication = inputData.YearOfPublication
	book.ImgUrlThumbnail = inputData.ImgUrlThumbnail
	book.ImgUrlCover = inputData.ImgUrlCover
	err = b.bookRepository.Update(book)
	if err != nil {
		return err
	}
	return nil
}

func (b *book) DeleteBook(id int) error {
	book, err := b.bookRepository.FindById(id)
	if err != nil {
		return err
	}
	if book.ID == 0 {
		return errors.New("data not found")
	}
	err = b.bookRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
