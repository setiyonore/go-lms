package helper

import (
	"go-lms/entities"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Auth interface {
	GenerateToken(User entities.User) (string, error)
}

type jwtAuth struct {
}

func NewAuth() *jwtAuth {
	return &jwtAuth{}
}

func (a *jwtAuth) GenerateToken(User entities.User) (string, error) {
	//create claims
	claims := jwt.MapClaims{
		"name":  User.Name,
		"email": User.Email,
		"id":    User.ID,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return t, err
	}
	return t, nil
}
