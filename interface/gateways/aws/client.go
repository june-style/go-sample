package aws

import (
	"github.com/june-style/go-sample/interface/gateways/aws/dynamodb"
	"github.com/june-style/go-sample/interface/gateways/aws/sqs"
)

type Client struct {
	DynamoDB *dynamodb.Client
	SQS      *sqs.Client
}
