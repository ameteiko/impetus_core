package main

import (
	"fmt"
	"log"

	"github.com/ameteiko/mindnet/src/domain/entity/dto"
	appdi "github.com/ameteiko/mindnet/src/internal/app/di"
	"github.com/ameteiko/mindnet/src/internal/config"
	"github.com/ameteiko/mindnet/src/internal/di"
)

//go:generate wire

func main() {
	cfg := config.NewConfig()

	dbSession, err := di.Neo4JSession(cfg.DB.URI, cfg.DB.User, cfg.DB.Pwd)
	if err != nil {
		log.Fatal("Cannot start application due to database connection error", err)
	}
	defer dbSession.Close()

	application := appdi.ProvideApp(dbSession)

	// Create a node.
	nodeDTO := dto.NewNode(
		dto.WithNodeTitle("Node Title"),
		dto.WithNodeSlug("node-title"),
		dto.WithNodeKind("artf"),
		dto.WithNodeContent("node content\ngoe\nhere"),
	)
	if _, err := application.CreateNode(nodeDTO); err != nil {
		fmt.Println("Error: ", err)
	}

	println("node was successfully created")
}
