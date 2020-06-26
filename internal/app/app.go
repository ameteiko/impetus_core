package app

import (
	"github.com/ameteiko/mindnet/domain/entity"
	"github.com/ameteiko/mindnet/domain/entity/dto"
	"github.com/ameteiko/mindnet/domain/value"
)

type (
	// TODO: Conceal all the interfaces for injected objects.
	EntityFactory interface {
		Node(dto.Node) (entity.Node, error)
	}
	NodeRepository interface {
		CreateNode(entity.Node) error
		GetNodeBySlug(value.Slug) (entity.Node, error)
	}
	NodeService interface {
		Relate(*entity.Node, entity.Relation)
	}

	App struct {
		entityFactory EntityFactory
		nodeRepo      NodeRepository
		nodeService   NodeService
	}
)

func New(entityFactory EntityFactory, nodeRepo NodeRepository, nodeService NodeService) App {
	return App{
		entityFactory: entityFactory,
		nodeRepo:      nodeRepo,
		nodeService:   nodeService,
	}
}

// CreateNode creates a new node in the application.
func (a App) CreateNode(n dto.Node) (entity.Node, error) {
	node, err := a.entityFactory.Node(n)
	if err != nil {
		return node, err
	}

	// Create a file entry if there was contentURI

	// Persist node to the database
	if err := a.nodeRepo.CreateNode(node); err != nil {
		return node, err
	}

	return node, nil
}

// func (a App) ConnectNodes(from string, to string, r dto.Relation) (rel value.Relation, err error) {
// 	fromSlug, _ := a.valueFactory.Slug(from)
// 	toSlug, _ := a.valueFactory.Slug(to)
// 	nodeFrom, err := a.nodeRepo.GetNodeBySlug(fromSlug)
// 	if err != nil {
// 		return rel, err
// 	}
//
// 	nodeTo, err := a.nodeRepo.GetNodeBySlug(toSlug)
// 	if err != nil {
// 		return rel, err
// 	}
//
// 	relation, err := a.valueFactory.Relation(r)
// 	if err != nil {
// 		return rel, err
// 	}
//
// 	rel = a.nodeService.Relate(nodeFrom, nodeTo, relation)
//
// 	// Persist node to the database
// 	if err := a.nodeRepo.CreateRelation(rel); err != nil {
// 		return rel, err
// 	}
//
// 	return rel, nil
// }
