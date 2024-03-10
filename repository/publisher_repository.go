package repository

import (
	"go-lms/entities"

	"gorm.io/gorm"
)

type Publisher interface {
	FindAll() ([]entities.Publisher, error)
	FindById(id int) (entities.Publisher, error)
	Save(publisher entities.Publisher) error
	Update(publisher entities.Publisher) error
	Delete(id int) error
}

type publisher struct {
	db *gorm.DB
}

func NewPublisherRepository(db *gorm.DB) *publisher {
	return &publisher{db: db}
}

func (p *publisher) FindAll() ([]entities.Publisher, error) {
	var publishers []entities.Publisher
	err := p.db.
		Select("id", "name").
		Find(&publishers).Error
	if err != nil {
		return publishers, err
	}
	return publishers, nil
}

func (p *publisher) FindById(id int) (entities.Publisher, error) {
	var publisher entities.Publisher
	err := p.db.Where("id=?", id).
		Select("id", "name").
		Find(&publisher).Error
	if err != nil {
		return publisher, err
	}
	return publisher, nil
}

func (p *publisher) Save(publisher entities.Publisher) error {
	err := p.db.Save(&publisher).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *publisher) Update(publisher entities.Publisher) error {
	err := p.db.Updates(&publisher).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *publisher) Delete(id int) error {
	var publisher entities.Publisher
	err := p.db.Where("id", id).First(&publisher).Error
	if err != nil {
		return err
	}
	err = p.db.Delete(&publisher).Error
	if err != nil {
		return err
	}
	return nil
}
