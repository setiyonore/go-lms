package service

import (
	"errors"
	"go-lms/entities"
	"go-lms/repository"
)

type Author interface {
	GetAuthor() ([]entities.Author, error)
	GetAuthorByID(ID int) (entities.Author, error)
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

func (a *author) GetAuthorByID(ID int) (entities.Author, error) {
	author, err := a.authorRepository.FindByID(ID)
	if err != nil {
		return author, err
	}
	if author.Id == 0 {
		err = errors.New("data not found")
		return author, err
	}
	return author, nil
}
