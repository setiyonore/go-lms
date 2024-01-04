package entities

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Id   int    `json:"id"`
	Name string `json:"name"`
}
