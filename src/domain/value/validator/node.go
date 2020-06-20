package validator

import (
	"errors"

	"github.com/ameteiko/mindnet/src/domain/value/dto"
)

var (
	ErrNodeTitleIsEmpty           = errors.New("node title is empty")
	ErrNodeSlugIsEmpty            = errors.New("node slug is empty")
)

type Node struct{}

func (v Node) Validate(n dto.Node) (errs []error) {
	if err := v.validateTitle(n.Title); err != nil  {
		errs = append(errs, err)
	}

	if err := v.validateSlug(n.Slug); err != nil {
		errs = append(errs, err)
	}

	return errs
}

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
