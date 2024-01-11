package service

import (
	"errors"
	"go-lms/entities"
	"go-lms/repository"
)

type Publisher interface {
	GetAll() ([]entities.Publisher, error)
	GetById(id int) (entities.Publisher, error)
}

type publiser struct {
	publisherRepository repository.Publisher
}

func NewService(publisherRepository repository.Publisher) *publiser {
	return &publiser{publisherRepository: publisherRepository}
}

func (p *publiser) GetAll() ([]entities.Publisher, error) {
	publisers, err := p.publisherRepository.FindAll()
	if err != nil {
		return publisers, err
	}
	return publisers, nil
}

func (p *publiser) GetById(id int) (entities.Publisher, error) {
	publiser, err := p.publisherRepository.FindById(id)
	if err != nil {
		return publiser, err
	}
	if publiser.Id == 0 {
		err = errors.New("data not found")
		return publiser, err
	}
	return publiser, nil
}
