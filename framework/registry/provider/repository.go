package provider

import (
	"github.com/google/wire"
	"github.com/june-style/go-sample/domain/entities"
	"github.com/june-style/go-sample/interface/repositories"
)

var RepositoryWireSet = wire.NewSet(
	repositories.NewRegisteredUserRepository,
	repositories.NewUserProfileRepository,
	repositories.NewUserSessionRepository,
	repositories.NewDynamoDBTx,
	wire.Struct(new(entities.Repository), "*"),
)
