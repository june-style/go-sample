package main

import (
	"context"
	"fmt"

	"github.com/june-style/go-sample/domain/derrors"
	"github.com/june-style/go-sample/interface/gateways/aws"
	"github.com/june-style/go-sample/interface/gateways/aws/dynamodb"
	"github.com/june-style/go-sample/interface/logs"
)

func NewService(client *aws.Client) *Service {
	return &Service{
		dbc: client.DynamoDB,
	}
}

type Service struct {
	dbc *dynamodb.Client
}

func (s *Service) MakeDB(ctx context.Context) error {
	for _, t := range tables {
		input := dynamodb.QueryInputData{
			Table: *t.TableName,
		}
		tableName, err := s.dbc.NewQuery(input).
			CreateTable(ctx, t)
		if err != nil {
			return derrors.Wrap(err)
		}
		logs.Info().Msg(fmt.Sprintf("made table '%s'", tableName))
	}
	return nil
}
