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
	ItemBooks         []itemBook     `gorm:"foreignkey:IdBook;references:ID"`
}

type itemBook struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	IdBook    uint           `json:"id_book" gorm:"index"`
	Isbn      string         `json:"isbn"`
	Status    int            `json:"status"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type AddBookInput struct {
	Name              string `json:"name" validate:"required,min=3"`
	Description       string `json:"description" validate:"required"`
	AuthorID          uint   `json:"author_id" validate:"required,min=1"`
	PublisherID       uint   `json:"publisher_id" validate:"required"`
	Isbn              string `json:"isbn" validate:"required"`
	YearOfPublication string `json:"year_of_publication" validate:"required"`
	ImgUrlThumbnail   string `json:"img_url_thumbnail"`
	ImgUrlCover       string `json:"img_url_cover"`
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
	bookFormatter.ItemBooks = book.ItemBooks
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

func FormatBookItem(ib itemBook) itemBook {
	itemBookFormatter := itemBook{}
	itemBookFormatter.ID = ib.ID
	itemBookFormatter.IdBook = ib.IdBook
	itemBookFormatter.Isbn = ib.Isbn
	itemBookFormatter.Status = ib.Status
	return itemBookFormatter
}
