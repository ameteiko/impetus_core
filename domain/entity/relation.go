package entity

import (
	"github.com/ameteiko/mindnet/domain/value"
)

type Relation struct {
	ID   value.ID
	Node Node
}
