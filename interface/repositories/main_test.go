package repositories_test

import (
	"context"

	"github.com/june-style/go-sample/domain/dconfig"
	"github.com/june-style/go-sample/interface/gateways/aws"
	"github.com/june-style/go-sample/interface/gateways/aws/dynamodb"
	"github.com/june-style/go-sample/interface/gateways/redis"
)

var (
	cfg         *dconfig.Config
	awsClient   *aws.Client
	redisClient *redis.Client
)

func init() {
	var (
		ctx = context.TODO()
		err error
	)
	cfg, err = dconfig.New("../../.env.template")
	if err != nil {
		panic(err.Error())
	}
	cfg.Aws.DynamoDBTable = "test"
	awsCfg, _ := aws.NewConfig(ctx, cfg)
	awsClient = &aws.Client{
		DynamoDB: dynamodb.NewClient(cfg, awsCfg),
	}
	redisClient, err = redis.NewClient(cfg)
	if err != nil {
		panic(err.Error())
	}
}
