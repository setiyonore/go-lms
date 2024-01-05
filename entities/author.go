package entities

import (
	"gorm.io/gorm"
	"time"
)

type Author struct {
	Id        int            `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
type AddAuthorInput struct {
	Name string `json:"name" validate:"required,min=3"`
}

func FormatAuthor(author Author) Author {
	authorFormatter := Author{}
	authorFormatter.Id = author.Id
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
