//go:build wireinject
// +build wireinject

package injector

import (
	"context"

	"github.com/google/wire"
	"github.com/june-style/go-sample/domain/dconfig"
	"github.com/june-style/go-sample/framework/registry/provider"
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
	Aws    *aws.Client
	Redis  *redis.Client
}
