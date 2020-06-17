package value

import (
	"github.com/ameteiko/mindnet/src/domain/value/dto"
	"github.com/ameteiko/mindnet/src/domain/value/sanitiser"
	"github.com/ameteiko/mindnet/src/domain/value/validator"
)

//go:generate stringer -type=NodeKind -linecomment -output=node_string.go
type NodeKind int

const (
	NodeArtefact NodeKind = iota // artefact
	NodeSnippet                  // snippet
)

// Node represents a MindNet node entity.
// Node entity is the basic data-storing entity.
type Node struct {
	kind NodeKind
	name string
}

//
func NewNode(n dto.Node) (node Node, err error) {
	if node, err = sanitiser.SanitiseNode(n); err != nil {
		return node, err
	}
	node = validator.ValidateNode(node)

	return Node{
		name: n.Slug,
		kind: n.Kind,
	}, nil
}

func (n Node) Name() string {
	return n.name
}
