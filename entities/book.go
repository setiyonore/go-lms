package entities

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID                uint           `json:"id" gorm:"primary_key"`
	Name              string         `json:"name"`
	Description       string         `json:"description"`
	PublisherID       uint           `json:"publisher_id"`
	AuthorID          uint           `json:"author_id"`
	Isbn              string         `json:"isbn"`
	YearOfPublication string         `json:"year_of_publication"`
	ImgUrlThumbnail   string         `json:"img_url_thumbnail"`
	ImgUrlCover       string         `json:"img_url_cover"`
	CreatedAt         time.Time      `json:"-"`
	UpdatedAt         time.Time      `json:"-"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`
	Publisher         Publisher      `gorm:"foreignkey:PublisherID"`
	Author            Author         `gorm:"foreignkey:AuthorID"`
}

func FormatBook(book Book) Book {
	bookFormatter := Book{}
	bookFormatter.ID = book.ID
	bookFormatter.Name = book.Name
	bookFormatter.Description = book.Description
	bookFormatter.PublisherID = book.PublisherID
	bookFormatter.AuthorID = book.AuthorID
	bookFormatter.Isbn = book.Isbn
	bookFormatter.YearOfPublication = book.YearOfPublication
	bookFormatter.ImgUrlThumbnail = book.ImgUrlThumbnail
	bookFormatter.ImgUrlCover = book.ImgUrlCover
	bookFormatter.Publisher.ID = book.PublisherID
	bookFormatter.Publisher.Name = book.Publisher.Name
	bookFormatter.Author.ID = book.AuthorID
	bookFormatter.Author.Name = book.Author.Name
	return bookFormatter
}

func FormatBooks(books []Book) []Book {
	booksFormatter := []Book{}
	for _, book := range books {
		bookFormatter := FormatBook(book)
		booksFormatter = append(booksFormatter, bookFormatter)
	}
	return booksFormatter
}
