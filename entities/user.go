package entities

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func FormatUser(user User) User {
	userFormatter := User{}
	userFormatter.ID = user.ID
	userFormatter.Name = user.Name
	userFormatter.Email = user.Email
	return userFormatter
}

func FormatUsers(users []User) []User {
	var usersFormatter []User
	for _, user := range users {
		userFormatter := FormatUser(user)
		usersFormatter = append(usersFormatter, userFormatter)
	}
	return usersFormatter
}
