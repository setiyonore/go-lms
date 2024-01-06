package helper

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func FormatterError(err validator.ValidationErrors) []string {
	var listError []string
	for _, fieldError := range err {
		listError = append(listError, fmt.Sprintf("error %s must be %s %s", fieldError.Field(), fieldError.Tag(), fieldError.Param()))
	}
	return listError
}
