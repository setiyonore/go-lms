package entities

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
type AddAuthorInput struct {
	Name string `json:"name" validate:"required,min=3"`
}

func FormatAuthor(author Author) Author {
	authorFormatter := Author{}
	authorFormatter.ID = author.ID
	authorFormatter.Name = author.Name
	return authorFormatter
}

func FormatAuthors(authors []Author) []Author {
	authorsFormatter := []Author{}
	for _, author := range authors {
		authorFormatter := FormatAuthor(author)
		authorsFormatter = append(authorsFormatter, authorFormatter)
	}
	return authorsFormatter
}
