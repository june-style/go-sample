//go:build wireinject
// +build wireinject

package injector

import (
	"context"

	"github.com/google/wire"
	"github.com/june-style/go-sample/application/usecases"
	"github.com/june-style/go-sample/domain/dconfig"
	"github.com/june-style/go-sample/domain/entities"
	"github.com/june-style/go-sample/framework/registry/provider"
	"github.com/june-style/go-sample/interface/controllers"
	"github.com/june-style/go-sample/interface/gateways/aws"
	"github.com/june-style/go-sample/interface/gateways/redis"
)

func InitDBClient(ctx context.Context, cfg *dconfig.Config) (*DBClient, error) {
	wire.Build(
		provider.DBClientWireSet,
		wire.Struct(new(aws.Client), "*"),
		wire.Struct(new(DBClient), "*"),
	)
	return nil, nil
}

type DBClient struct {
	Config *dconfig.Config

	Aws   *aws.Client
	Redis *redis.Client
}

func InitAPI(dbClient *DBClient) (*API, error) {
	wire.Build(
		provider.APIWireSet,
		wire.Struct(new(API), "*"),
		wire.FieldsOf(new(*DBClient), "Config"),
		wire.FieldsOf(new(*DBClient), "AWs"),
		wire.FieldsOf(new(*DBClient), "Redis"),
	)
	return nil, nil
}

type API struct {
	Config *dconfig.Config

	Controller *controllers.Controller
	Repository *entities.Repository
	UseCase    *usecases.UseCase
}
