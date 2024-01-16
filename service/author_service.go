package service

import (
	"errors"
	"go-lms/entities"
	"go-lms/repository"
)

type Author interface {
	GetAuthor() ([]entities.Author, error)
	GetAuthorByID(ID int) (entities.Author, error)
	AddAuhtor(input entities.AddAuthorInput) error
	UpdateAuthor(inputID int, inputData entities.AddAuthorInput) error
	DeleteAuthor(ID int) error
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
	if author.ID == 0 {
		err = errors.New("data not found")
		return author, err
	}
	return author, nil
}

func (a *author) AddAuhtor(input entities.AddAuthorInput) error {
	author := entities.Author{}
	author.Name = input.Name
	err := a.authorRepository.Save(author)
	if err != nil {
		return err
	}
	return nil

}

func (a *author) UpdateAuthor(inputID int, inputData entities.AddAuthorInput) error {
	author, err := a.authorRepository.FindByID(inputID)
	if err != nil {
		return err
	}
	if author.ID == 0 {
		return errors.New("data not found")
	}
	author.Name = inputData.Name
	err = a.authorRepository.Update(author)
	if err != nil {
		return err
	}
	return nil
}

func (a *author) DeleteAuthor(ID int) error {
	author, err := a.authorRepository.FindByID(ID)
	if err != nil {
		return err
	}
	if author.ID == 0 {
		return errors.New("data not found")
	}
	err = a.authorRepository.Delete(ID)
	if err != nil {
		return err
	}
	return nil
}
