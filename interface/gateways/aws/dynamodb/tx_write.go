package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/june-style/go-sample/domain/derrors"
)

func NewTxWrite(db *dynamodb.Client) *txWrite {
	return &txWrite{
		db:           db,
		txWriteInput: &dynamodb.TransactWriteItemsInput{},
	}
}

type txWrite struct {
	db           *dynamodb.Client
	txWriteInput *dynamodb.TransactWriteItemsInput
}

func (t *txWrite) RunWithContext(ctx context.Context) error {
	if _, err := t.db.TransactWriteItems(ctx, t.txWriteInput); err != nil {
		return derrors.Wrap(err)
	}
	return nil
}

func (t *txWrite) PutItem(ctx context.Context, item *types.Put) {
	t.txWriteInput.TransactItems = append(t.txWriteInput.TransactItems, types.TransactWriteItem{
		Put: item,
	})
}

func (t *txWrite) UpdateItem(ctx context.Context, item *types.Update) {
	t.txWriteInput.TransactItems = append(t.txWriteInput.TransactItems, types.TransactWriteItem{
		Update: item,
	})
}

func (t *txWrite) DeleteItem(ctx context.Context, item *types.Delete) {
	t.txWriteInput.TransactItems = append(t.txWriteInput.TransactItems, types.TransactWriteItem{
		Delete: item,
	})
}
