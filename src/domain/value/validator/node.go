package validator

import (
	"errors"

	"github.com/ameteiko/mindnet/src/domain/value/dto"
)

var (
	ErrNodeTitleIsEmpty           = errors.New("node title is empty")
	ErrNodeSlugIsEmpty            = errors.New("node slug is empty")
	ErrNodeSlugMustStartWithAWord = errors.New("node slug must start with an alphanumeric character")
)

type Node struct{}

func (v Node) Validate(n dto.Node) error {
	return nil
}

// Slug cannot start or and with - character.
func (_ Node) validateSlug(s string) error {
	if s == "" {
		return ErrNodeSlugIsEmpty
	}

	return nil
}

func (_ Node) validateTitle(t string) error {
	if t == "" {
		return ErrNodeTitleIsEmpty
	}

	return nil
}
