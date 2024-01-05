package entities

import "time"

type Author struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}
type AddAuthorInput struct {
	Name string `json:"name"`
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
