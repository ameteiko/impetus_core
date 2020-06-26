package aggregate

import (
	"github.com/ameteiko/mindnet/src/domain/entity"
)

type Node struct {
	root entity.Node
	relations []entity.Relation
}
