package service

import (
	"errors"
	"go-lms/entities"
	"go-lms/repository"
)

type Publisher interface {
	GetAll() ([]entities.Publisher, error)
	GetById(id int) (entities.Publisher, error)
	AddPublisher(input entities.AddPublisherInput) error
	UpdatePublisher(inputID int, inputData entities.AddPublisherInput) error
	DeletePublisher(id int) error
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

func (p *publiser) AddPublisher(input entities.AddPublisherInput) error {
	publiser := entities.Publisher{}
	publiser.Name = input.Name
	err := p.publisherRepository.Save(publiser)
	if err != nil {
		return err
	}
	return nil
}

func (p *publiser) UpdatePublisher(inputID int, inputData entities.AddPublisherInput) error {
	publisher, err := p.publisherRepository.FindById(inputID)
	if err != nil {
		return err
	}
	if publisher.Id == 0 {
		return errors.New("data not found")
	}
	publisher.Name = inputData.Name
	err = p.publisherRepository.Update(publisher)
	if err != nil {
		return err
	}
	return nil
}

func (p *publiser) DeletePublisher(id int) error {
	publisher, err := p.publisherRepository.FindById(id)
	if err != nil {
		return err
	}
	if publisher.Id == 0 {
		return errors.New("data not found")
	}
	err = p.publisherRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
