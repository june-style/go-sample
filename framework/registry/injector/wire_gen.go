// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"context"
	"github.com/june-style/go-sample/domain/dconfig"
	"github.com/june-style/go-sample/interface/gateways/aws"
	"github.com/june-style/go-sample/interface/gateways/aws/dynamodb"
	"github.com/june-style/go-sample/interface/gateways/aws/sqs"
	"github.com/june-style/go-sample/interface/gateways/redis"
)

// Injectors from wire.go:

func InitDBClient(ctx context.Context, cfg *dconfig.Config) (*DBClient, error) {
	config, err := aws.NewConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}
	client := dynamodb.NewClient(cfg, config)
	sqsClient := sqs.NewClient(cfg, config)
	awsClient := &aws.Client{
		DynamoDB: client,
		SQS:      sqsClient,
	}
	redisClient, err := redis.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	dbClient := &DBClient{
		Config: cfg,
		Aws:    awsClient,
		Redis:  redisClient,
	}
	return dbClient, nil
}

// wire.go:

type DBClient struct {
	Config *dconfig.Config
	Aws    *aws.Client
	Redis  *redis.Client
}
