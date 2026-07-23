package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field string
}

func (e ValidationError) Error() string {
	return e.Field + " is invalid"
}

func validateUser() error {

	err := ValidationError{
		Field: "email",
	}

	return fmt.Errorf("validation failed: %w", err)
}

func main() {

	err := validateUser()

	var validationErr ValidationError

	if errors.As(err, &validationErr) {
		fmt.Println("Validation error found")
		fmt.Println("Invalid field:", validationErr.Field)
	}
}