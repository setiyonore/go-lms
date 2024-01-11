package service

import (
	"go-lms/entities"
	"go-lms/repository"
)

type Publisher interface {
	GetAll() ([]entities.Publisher, error)
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
