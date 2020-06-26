package validator

import (
	"errors"
)

var (
	ErrIDIsEmpty    = errors.New("ID is empty")
	ErrSlugIsEmpty  = errors.New("slug is empty")
	ErrTitleIsEmpty = errors.New("title is empty")
)

type Validator struct{}

func NewValidator() Validator {
	return Validator{}
}

func (_ Validator) ValidateID(id string) error {
	if id == "" {
		return ErrIDIsEmpty
	}

	return nil
}

func (_ Validator) ValidateSlug(s string) error {
	if s == "" {
		return ErrSlugIsEmpty
	}

	return nil
}

func (_ Validator) ValidateTitle(t string) error {
	if t == "" {
		return ErrTitleIsEmpty
	}

	return nil
}
