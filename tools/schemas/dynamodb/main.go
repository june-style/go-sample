package main

import (
	"context"

	"github.com/june-style/go-sample/domain/dconfig"
	"github.com/june-style/go-sample/framework/registry/injector"
	"github.com/june-style/go-sample/interface/logs"
)

func main() {
	if err = svc.MakeDB(ctx); err != nil {
		logs.Fatal().SetError(err).Send()
	}
	logs.Info().Msg("...done!")
}

var (
	ctx context.Context
	cfg *dconfig.Config
	dbc *injector.DBClient
	svc *Service
	err error
)

func init() {
	ctx = context.Background()
	if cfg, err = dconfig.New(); err != nil {
		logs.Fatal().SetError(err).Msg("Stopped")
	}
	if dbc, err = injector.InitDBClient(ctx, cfg); err != nil {
		logs.Fatal().SetError(err).Msg("Stopped")
	}
	svc = NewService(dbc.Aws)
}
