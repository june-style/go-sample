package provider

import (
	"github.com/google/wire"
	"github.com/june-style/go-sample/interface/gateways/aws"
	"github.com/june-style/go-sample/interface/gateways/aws/dynamodb"
	"github.com/june-style/go-sample/interface/gateways/aws/sqs"
	"github.com/june-style/go-sample/interface/gateways/redis"
)

var DBClientWireSet = wire.NewSet(
	aws.NewConfig,
	dynamodb.NewClient,
	sqs.NewClient,
	redis.NewClient,
)
