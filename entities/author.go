package entities

type Author struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
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
