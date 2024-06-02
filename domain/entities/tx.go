package entities

import "context"

type DynamoDB interface {
	TxWrite(ctx context.Context, fnc func(ctx context.Context) error) error
}
