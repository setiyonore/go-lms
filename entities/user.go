package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	// Role      int            `json:"role"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
type GetUserByEmailInput struct {
	Email string `json:"email" validate:"required,email"`
}

type AddUserInput struct {
	Name     string `json:"name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	// Role     int    `json:"role"`
}
type EditUserInput struct {
	Name     string `json:"name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
	// Role     int    `json:"role"`
}

func FormatUser(user User) User {
	userFormatter := User{}
	userFormatter.ID = user.ID
	userFormatter.Name = user.Name
	userFormatter.Email = user.Email
	// userFormatter.Role = user.Role
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
