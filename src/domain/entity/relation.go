package entity

import (
	"github.com/ameteiko/mindnet/src/domain/value"
)

type Relation struct {
	ID   value.ID
	Node Node
}
