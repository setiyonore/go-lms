package repository

import "go-lms/entities"

type BookBorrowing interface {
	GetAll() ([]entities.BookBorrowings, error)
}
