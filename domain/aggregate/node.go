package aggregate

import (
	"github.com/ameteiko/mindnet/domain/entity"
)

type Node struct {
	root      entity.Node
	relations []entity.Relation
}
