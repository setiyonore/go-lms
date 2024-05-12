package entities

type ItemBook struct {
	ID     int
	IdBook int
	Isbn   string
	Status int
}

type AddItemBookInput struct {
	IdBook int    `json:"id_book" validate:"required"`
	Isbn   string `json:"isbn" validate:"required"`
	Status string `json:"status" validate:"required"`
}

func FormatItemBook(itemBook ItemBook) ItemBook {
	formatter := ItemBook{}

	formatter.ID = itemBook.ID
	formatter.IdBook = itemBook.IdBook
	formatter.Isbn = itemBook.Isbn
	formatter.Status = itemBook.Status
	return formatter
}

func FormatItemBooks(itemBooks []ItemBook) []ItemBook {
	formatters := []ItemBook{}
	for _, bookitem := range itemBooks {
		formatter := FormatItemBook(bookitem)
		formatters = append(formatters, formatter)
	}
	return formatters
}
