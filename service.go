package main

import (
	"errors"
	"strings"
)

// ErrEmpty is an error object representing empty string errors
var ErrEmpty = errors.New("Empty string")

// StringService is an interface that models the String Service and contains the business logic
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

// stringService is a struct that implements the StringService
type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}
