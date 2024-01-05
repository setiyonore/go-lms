package service

import (
	"go-lms/entities"
	"go-lms/repository"
)

type Author interface {
	GetAuthor() ([]entities.Author, error)
}

type author struct {
	authorRepository repository.Author
}

func NewAuthor(authorRepository repository.Author) *author {
	return &author{authorRepository}
}

func (a *author) GetAuthor() ([]entities.Author, error) {
	authors, err := a.authorRepository.FindAll()
	if err != nil {
		return authors, err
	}
	return authors, nil
}
