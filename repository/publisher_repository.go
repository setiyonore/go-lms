package repository

import (
	"go-lms/entities"

	"gorm.io/gorm"
)

type Publisher interface {
	FindAll() ([]entities.Publisher, error)
}

type publisher struct {
	db *gorm.DB
}

func NewPublisherRepository(db *gorm.DB) *publisher {
	return &publisher{db: db}
}

func (p *publisher) FindAll() ([]entities.Publisher, error) {
	var publishers []entities.Publisher
	err := p.db.Find(&publishers).Error
	if err != nil {
		return publishers, err
	}
	return publishers, nil
}
