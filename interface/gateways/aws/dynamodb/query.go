package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/june-style/go-sample/domain/derrors"
)

func (c *Client) NewQuery(q QueryInputData) Query {
	return &queryImpl{
		db:    c.GetDB(),
		table: c.GetTableFullName(q.GetTableName()),
	}
}

type Query interface {
	Get(ctx context.Context, keys Keys, orm any) error
	Put(ctx context.Context, orm any) error
	Post(ctx context.Context, keys Keys, values map[string]any) error
	Del(ctx context.Context, keys Keys) error

	CreateTable(ctx context.Context, params dynamodb.CreateTableInput) (string, error)
}

type queryImpl struct {
	db    *dynamodb.Client
	table string
}

func (q *queryImpl) Get(ctx context.Context, keys Keys, orm any) error {
	res, err := q.db.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(q.table),
		Key:       keys.MarshalAttributeValues(),
	})
	if err != nil {
		return derrors.Wrap(err)
	}
	if err := attributevalue.UnmarshalMap(res.Item, orm); err != nil {
		return derrors.Wrap(err)
	}
	return nil
}

func (q *queryImpl) Put(ctx context.Context, orm any) error {
	av, err := attributevalue.MarshalMap(orm)
	if err != nil {
		return derrors.Wrap(err)
	}
	if txWrite := GetTxWriteByContext(ctx); txWrite != nil {
		txWrite.PutItem(ctx, &types.Put{
			TableName: aws.String(q.table),
			Item:      av,
		})
	} else {
		if _, err := q.db.PutItem(ctx, &dynamodb.PutItemInput{
			TableName: aws.String(q.table),
			Item:      av,
		}); err != nil {
			return derrors.Wrap(err)
		}
	}
	return nil
}

func (q *queryImpl) Post(ctx context.Context, keys Keys, values map[string]any) error {
	update := expression.UpdateBuilder{}
	for k, v := range values {
		update = update.Set(expression.Name(k), expression.Value(v))
	}
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		return derrors.Wrap(err)
	}
	if txWrite := GetTxWriteByContext(ctx); txWrite != nil {
		txWrite.UpdateItem(ctx, &types.Update{
			TableName:                 aws.String(q.table),
			Key:                       keys.MarshalAttributeValues(),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			UpdateExpression:          expr.Update(),
		})
	} else {
		if _, err = q.db.UpdateItem(ctx, &dynamodb.UpdateItemInput{
			TableName:                 aws.String(q.table),
			Key:                       keys.MarshalAttributeValues(),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			UpdateExpression:          expr.Update(),
		}); err != nil {
			return derrors.Wrap(err)
		}
	}
	return nil
}

func (q *queryImpl) Del(ctx context.Context, keys Keys) error {
	if txWrite := GetTxWriteByContext(ctx); txWrite != nil {
		txWrite.DeleteItem(ctx, &types.Delete{
			TableName: aws.String(q.table),
			Key:       keys.MarshalAttributeValues(),
		})
	} else {
		if _, err := q.db.DeleteItem(ctx, &dynamodb.DeleteItemInput{
			TableName: aws.String(q.table),
			Key:       keys.MarshalAttributeValues(),
		}); err != nil {
			return derrors.Wrap(err)
		}
	}
	return nil
}

func (q *queryImpl) CreateTable(ctx context.Context, params dynamodb.CreateTableInput) (string, error) {
	params.TableName = aws.String(q.table)
	if _, err := q.db.CreateTable(ctx, &params); err != nil {
		return "", derrors.Wrap(err)
	}
	return q.table, nil
}
