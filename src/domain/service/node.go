package service

import (
	"github.com/ameteiko/mindnet/src/domain/entity"
)

type Node struct {}

func NewNode() Node {
	return Node{}
}

func (_ Node) Relate(n *entity.Node, r entity.Relation) {

}
