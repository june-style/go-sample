package dconfig

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/june-style/go-sample/domain/derrors"
)

type Config struct {
	App   App
	Aws   Aws
	Grpc  Grpc
	Redis Redis
	Sys   Sys
}

func New(dotenv ...string) (*Config, error) {
	if err := godotenv.Load(dotenv...); err != nil {
		return nil, derrors.Wrap(err)
	}
	cfg := &Config{}

	if err := env.Parse(&cfg.App); err != nil {
		return nil, derrors.Wrap(err)
	}
	if err := env.Parse(&cfg.Aws); err != nil {
		return nil, derrors.Wrap(err)
	}
	if err := env.Parse(&cfg.Grpc); err != nil {
		return nil, derrors.Wrap(err)
	}
	if err := env.Parse(&cfg.Redis); err != nil {
		return nil, derrors.Wrap(err)
	}
	if err := env.Parse(&cfg.Sys); err != nil {
		return nil, derrors.Wrap(err)
	}
	return cfg, nil
}
