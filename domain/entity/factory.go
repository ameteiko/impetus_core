package entity

import (
	"github.com/ameteiko/mindnet/domain/entity/dto"
	"github.com/ameteiko/mindnet/domain/value"
)

type (
	// TODO: Conceal all the interfaces for injected objects.
	IDGenerator interface{ Generate() string }

	ValueFactory interface {
		ID(string) (value.ID, error)
		Title(string) (value.Title, error)
		Slug(string, string) (value.Slug, error)
	}

	Factory struct {
		valueFactory ValueFactory
		idGenerator  IDGenerator
	}
)

func NewFactory(valueFactory ValueFactory, idGenerator IDGenerator) Factory {
	return Factory{
		valueFactory: valueFactory,
		idGenerator:  idGenerator,
	}
}

func (f Factory) Node(n dto.Node) (Node, error) {
	var (
		id    value.ID
		title value.Title
		slug  value.Slug
		err   error
	)

	if id, err = f.valueFactory.ID(f.idGenerator.Generate()); err != nil {
		return Node{}, err
	}

	if title, err = f.valueFactory.Title(n.Title); err != nil {
		return Node{}, nil
	}

	if slug, err = f.valueFactory.Slug(n.Slug, n.Title); err != nil {
		return Node{}, nil
	}

	return Node{
		ID:    id,
		Title: title,
		Slug:  slug,
	}, nil
}

// Relation is a multiplicity association for the Node.
func (f Factory) Relation(to Node, r dto.Relation) (rel Relation, err error) {

	return rel, nil
}
