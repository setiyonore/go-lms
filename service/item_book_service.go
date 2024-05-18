package service

import (
	"go-lms/entities"
	"go-lms/repository"
)

type ItemBook interface {
	GetItemBook() ([]entities.ItemBook, error)
	GetItemBookById(id int) (entities.ItemBook, error)
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
