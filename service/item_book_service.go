package service

import (
	"errors"
	"go-lms/entities"
	"go-lms/repository"
)

type ItemBook interface {
	GetItemBook() ([]entities.ItemBook, error)
	GetItemBookById(id int) (entities.ItemBook, error)
	AddItemBook(input entities.AddItemBookInput) error
	UpdateItemBook(inputId int, input entities.AddItemBookInput) error
	UpdateStatusItemBook(id int, status int) error
	DeleteItemBook(id int) error
}
type itemBook struct {
	itemBookRepository repository.ItemBook
}

func NewItemBook(itemBookRepository repository.ItemBook) *itemBook {
	return &itemBook{itemBookRepository: itemBookRepository}
}

func (s *itemBook) GetItemBook() ([]entities.ItemBook, error) {
	itemBooks, err := s.itemBookRepository.FindAll()
	if err != nil {
		return itemBooks, err
	}
	return itemBooks, nil
}

func (s *itemBook) GetItemBookById(id int) (entities.ItemBook, error) {
	itemBook, err := s.itemBookRepository.FIndById(id)
	if err != nil {
		return itemBook, err
	}
	return itemBook, nil
}

func (s *itemBook) AddItemBook(input entities.AddItemBookInput) error {
	itemBook := entities.ItemBook{}
	itemBook.IdBook = input.IdBook
	itemBook.Isbn = input.Isbn
	itemBook.Status = input.Status
	err := s.itemBookRepository.Save(itemBook)
	if err != nil {
		return err
	}
	return nil
}

func (s *itemBook) UpdateItemBook(inputId int, input entities.AddItemBookInput) error {
	itemBook, err := s.itemBookRepository.FIndById(inputId)
	if err != nil {
		return err
	}
	if itemBook.ID == 0 {
		err = errors.New("data not found")
		return err
	}
	itemBook = entities.ItemBook{}
	itemBook.ID = inputId
	itemBook.IdBook = input.IdBook
	itemBook.Isbn = input.Isbn
	itemBook.Status = input.Status
	err = s.itemBookRepository.Update(itemBook)
	if err != nil {
		return err
	}
	return nil
}

func (s *itemBook) UpdateStatusItemBook(id int, status int) error {
	itemBook, err := s.itemBookRepository.FIndById(id)
	if err != nil {
		return err
	}
	if itemBook.ID == 0 {
		err = errors.New("data not found")
		return err
	}
	err = s.itemBookRepository.UpdateStatus(id, status)
	if err != nil {
		return err
	}
	return nil
}

func (s *itemBook) DeleteItemBook(id int) error {
	itemBook, err := s.itemBookRepository.FIndById(id)
	if err != nil {
		return err
	}
	if itemBook.ID == 0 {
		return errors.New("data not found")
	}
	err = s.itemBookRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
