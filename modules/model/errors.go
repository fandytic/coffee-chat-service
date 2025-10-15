package model

import "errors"

var (
	ErrCustomerNotFound = errors.New("customer not found")
	ErrOrderNotFound    = errors.New("order not found")
)

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}
