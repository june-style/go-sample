package interactors_test

import (
	"context"

	"github.com/june-style/go-sample/domain/dconfig"
	"github.com/june-style/go-sample/domain/entities"
)

var cfg *dconfig.Config

func init() {
	var err error
	cfg, err = dconfig.New("../../.env.template")
	if err != nil {
		panic(err.Error())
	}
}

func NewMockDynamoDB() entities.DynamoDB {
	return &mockDynamoDB{}
}

type mockDynamoDB struct{}

func (m mockDynamoDB) TxWrite(ctx context.Context, fnc func(context.Context) error) error {
	return fnc(ctx)
}
