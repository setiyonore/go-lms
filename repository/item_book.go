package repository

import "go-lms/entities"

type ItemBook interface {
	FindAll() ([]entities.ItemBook, error)
}
