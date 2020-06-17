package app

import (
	"github.com/ameteiko/mindnet/src/domain/value"
)

type App struct {
}

// CreateNode creates a new node in the application.
// CreateNode receives a list of functional options.
func (a App) CreateNode(name string) (node value.Node, err error) {
	return value.NewNode(value.WithName(name))
}