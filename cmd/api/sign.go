package main

import (
	"context"

	"github.com/june-style/go-sample/domain/dconfig"
	"github.com/june-style/go-sample/domain/dcontext"
	"github.com/june-style/go-sample/domain/entities"
	"github.com/june-style/go-sample/framework/protocol/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func newSign(cfg *dconfig.Config, conn *grpc.ClientConn) *sign {
	return &sign{
		config: cfg,
		conn:   conn,
	}
}

type sign struct {
	config *dconfig.Config
	conn   *grpc.ClientConn
}

func (s *sign) in(ctx context.Context, accessKey string) (*pb.SignInResponse, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
		dcontext.HeaderApplicationKey.String(): s.config.App.Key,
		dcontext.HeaderAccessKey.String():      accessKey,
		dcontext.HeaderRequestID.String():      entities.GenULID(),
	}))

	sign := pb.NewSignServiceClient(s.conn)

	if res, err := sign.In(ctx, &pb.SignInRequest{}); err != nil {
		return nil, err
	} else {
		return res, nil
	}
}

func (s *sign) up(ctx context.Context, name string) (*pb.SignUpResponse, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
		dcontext.HeaderApplicationKey.String(): s.config.App.Key,
		dcontext.HeaderRequestID.String():      entities.GenULID(),
	}))

	sign := pb.NewSignServiceClient(s.conn)

	if res, err := sign.Up(ctx, &pb.SignUpRequest{
		Sign: name,
	}); err != nil {
		return nil, err
	} else {
		return res, nil
	}
}
