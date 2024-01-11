package entities

import (
	"time"

	"gorm.io/gorm"
)

type Publisher struct {
	Id        int            `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"-"`
	UPdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func FormatterPublisher(publisher Publisher) Publisher {
	publisherFormatter := Publisher{}
	publisherFormatter.Id = publisher.Id
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
