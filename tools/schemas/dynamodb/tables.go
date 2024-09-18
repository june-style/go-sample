package main

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var tables = []dynamodb.CreateTableInput{
	data,
	test,
}

var data = dynamodb.CreateTableInput{
	TableName: aws.String("data"),
	KeySchema: []types.KeySchemaElement{
		{
			AttributeName: aws.String("pkey"),
			KeyType:       types.KeyTypeHash,
		},
		{
			AttributeName: aws.String("skey"),
			KeyType:       types.KeyTypeRange,
		},
	},
	AttributeDefinitions: []types.AttributeDefinition{
		{
			AttributeName: aws.String("pkey"),
			AttributeType: types.ScalarAttributeTypeS,
		},
		{
			AttributeName: aws.String("skey"),
			AttributeType: types.ScalarAttributeTypeS,
		},
	},
	BillingMode: types.BillingModePayPerRequest,
}

var test = dynamodb.CreateTableInput{
	TableName: aws.String("test"),
	KeySchema: []types.KeySchemaElement{
		{
			AttributeName: aws.String("pkey"),
			KeyType:       types.KeyTypeHash,
		},
		{
			AttributeName: aws.String("skey"),
			KeyType:       types.KeyTypeRange,
		},
	},
	AttributeDefinitions: []types.AttributeDefinition{
		{
			AttributeName: aws.String("pkey"),
			AttributeType: types.ScalarAttributeTypeS,
		},
		{
			AttributeName: aws.String("skey"),
			AttributeType: types.ScalarAttributeTypeS,
		},
	},
	BillingMode: types.BillingModePayPerRequest,
}
