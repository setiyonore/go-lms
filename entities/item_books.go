package entities

type BookItems struct {
	Id     int
	IdBook int
	Isbn   string
	Status string
}

type AddBookItemsInput struct {
	Isbn   string `json:"isbn" validate:"required"`
	Status string `json:"status" validate:"required"`
}
