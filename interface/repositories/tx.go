package repositories

import (
	"context"

	"github.com/june-style/go-sample/domain/derrors"
	"github.com/june-style/go-sample/domain/entities"
	"github.com/june-style/go-sample/interface/gateways/aws"
	"github.com/june-style/go-sample/interface/gateways/aws/dynamodb"
)

func NewDynamoDBTx(awsClient *aws.Client) entities.DynamoDB {
	return &dynamoDBTxImpl{
		db: awsClient.DynamoDB,
	}
}

type dynamoDBTxImpl struct {
	db *dynamodb.Client
}

func (t *dynamoDBTxImpl) TxWrite(ctx context.Context, fnc func(context.Context) error) error {
	txWrite := dynamodb.GetTxWriteByContext(ctx)
	if txWrite == nil {
		txWrite = dynamodb.NewTxWrite(t.db.GetDB())
		ctx = dynamodb.SetTxWriteByContext(ctx, txWrite)
	}

	if err := fnc(ctx); err != nil {
		return derrors.Wrap(err)
	}

	if err := txWrite.RunWithContext(ctx); err != nil {
		return derrors.Wrap(err)
	}

	return nil
}
