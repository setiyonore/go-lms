package entities

import (
	"time"

	"gorm.io/gorm"
)

type Publisher struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
type AddPublisherInput struct {
	Name string `json:"name" validate:"required,min=3"`
}

func FormatterPublisher(publisher Publisher) Publisher {
	publisherFormatter := Publisher{}
	publisherFormatter.ID = publisher.ID
	publisherFormatter.Name = publisher.Name
	return publisherFormatter
}

func FormatterPublishers(publishers []Publisher) []Publisher {
	publishersFormatter := []Publisher{}
	for _, publiser := range publishers {
		publisherFormatter := FormatterPublisher(publiser)
		publishersFormatter = append(publishersFormatter, publisherFormatter)
	}
	return publishersFormatter
}
