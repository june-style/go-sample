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

func newHome(cfg *dconfig.Config, conn *grpc.ClientConn) *home {
	return &home{
		config: cfg,
		conn:   conn,
	}
}

type home struct {
	config *dconfig.Config
	conn   *grpc.ClientConn
}

func (h *home) get(ctx context.Context, accessKey, sessionID string) (*pb.HomeGetResponse, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
		dcontext.HeaderApplicationKey.String(): h.config.App.Key,
		dcontext.HeaderAccessKey.String():      accessKey,
		dcontext.HeaderSessionID.String():      sessionID,
		dcontext.HeaderRequestID.String():      entities.GenULID(),
	}))

	home := pb.NewHomeServiceClient(h.conn)

	if res, err := home.Get(ctx, &pb.HomeGetRequest{}); err != nil {
		return nil, err
	} else {
		return res, nil
	}
}
