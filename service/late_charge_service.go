package service

import (
	"go-lms/entities"
	"go-lms/repository"
)

type LateCharge interface {
	GetLateCharge() ([]entities.LateCharge, error)
	GetLateChargeById(id int) (entities.LateCharge, error)
}

type latecharge struct {
	latechargeRepository repository.LateCharge
}

func NewLateCharge(latechargeRepository repository.LateCharge) *latecharge {
	return &latecharge{latechargeRepository: latechargeRepository}
}

func (s *latecharge) GetLateCharge() ([]entities.LateCharge, error) {
	latecharge, err := s.latechargeRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return latecharge, nil
}

func (s *latecharge) GetLateChargeById(id int) (entities.LateCharge, error) {
	latecharge, err := s.latechargeRepository.FindById(id)
	if err != nil {
		return latecharge, err
	}
	return latecharge, nil
}
