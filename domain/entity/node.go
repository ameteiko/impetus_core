package entity

import (
	"github.com/ameteiko/mindnet/domain/value"
)

type Node struct {
	ID    value.ID
	Title value.Title
	Slug  value.Slug
}
