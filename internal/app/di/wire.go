// +build wireinject

package di

import (
	"github.com/google/wire"
	_neo4j "github.com/neo4j/neo4j-go-driver/neo4j"

	"github.com/ameteiko/mindnet/domain/entity"
	"github.com/ameteiko/mindnet/domain/service"
	"github.com/ameteiko/mindnet/domain/value"
	"github.com/ameteiko/mindnet/domain/value/sanitiser"
	"github.com/ameteiko/mindnet/domain/value/validator"
	"github.com/ameteiko/mindnet/internal/app"
	"github.com/ameteiko/mindnet/internal/platform/neo4j"
	"github.com/ameteiko/mindnet/internal/platform/services"
)

func provideValueSanitiser() (s sanitiser.Sanitiser) {
	wire.Build(wire.Struct(new(sanitiser.Sanitiser)))
	return s
}

func provideValueValidator() (v validator.Validator) {
	wire.Build(wire.Struct(new(validator.Validator)))
	return v
}

func provideIDGenerator() (idGenerator services.IDGenerator) {
	wire.Build(wire.Struct(new(services.IDGenerator)))
	return idGenerator
}

func provideNodeRepository(dbSession _neo4j.Session) (nodeRepository *neo4j.DB) {
	wire.Build(neo4j.NewDB)
	return nodeRepository
}

func provideNodeService() (nodeService service.Node) {
	wire.Build(wire.Struct(new(service.Node)))
	return nodeService
}

func provideValueFactory() (valueFactory value.Factory) {
	wire.Build(wire.NewSet(
		value.NewFactory,
		provideValueSanitiser,
		provideValueValidator,
		wire.Bind(new(value.Sanitiser), new(sanitiser.Sanitiser)),
		wire.Bind(new(value.Validator), new(validator.Validator)),
	))
	return valueFactory
}

func provideEntityFactory() (entityFactory entity.Factory) {
	wire.Build(wire.NewSet(
		entity.NewFactory,
		provideValueFactory,
		provideIDGenerator,
		wire.Bind(new(entity.ValueFactory), new(value.Factory)),
		wire.Bind(new(entity.IDGenerator), new(services.IDGenerator)),
	))
	return entityFactory
}

func ProvideApp(dbSession _neo4j.Session) (application app.App) {
	wire.Build(
		app.New,
		provideEntityFactory,
		provideNodeRepository,
		provideNodeService,
		wire.Bind(new(app.EntityFactory), new(entity.Factory)),
		wire.Bind(new(app.NodeRepository), new(*neo4j.DB)),
		wire.Bind(new(app.NodeService), new(service.Node)),
	)

	return application
}
