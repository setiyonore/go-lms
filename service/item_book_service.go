package service

import (
	"go-lms/entities"
	"go-lms/repository"
)

type ItemBook interface {
	GetItemBook() ([]entities.ItemBook, error)
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
