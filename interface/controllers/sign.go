package controllers

import (
	"context"

	"github.com/june-style/go-sample/application/usecases"
	"github.com/june-style/go-sample/domain/dcontext"
	"github.com/june-style/go-sample/domain/derrors"
	"github.com/june-style/go-sample/framework/protocol/pb"
)

func NewSign(usecase *usecases.UseCase) Sign {
	return &signImpl{
		usecases: usecase,
	}
}

type Sign interface {
	pb.SignServiceServer
}

type signImpl struct {
	pb.UnimplementedSignServiceServer
	usecases *usecases.UseCase
}

func (c *signImpl) In(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	output, err := c.usecases.Sign.In(ctx, usecases.SignInInputData{
		UserID: dcontext.GetAuthenticatedUserID(ctx),
	})
	if err != nil {
		return nil, derrors.Wrap(err)
	}
	return &pb.SignInResponse{
		SessionId: output.UserSession.SessionID(),
	}, nil
}

func (c *signImpl) Up(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	output, err := c.usecases.Sign.Up(ctx, usecases.SignUpInputData{
		Sign: req.Sign,
	})
	if err != nil {
		return nil, derrors.Wrap(err)
	}
	return &pb.SignUpResponse{
		AccessKey: output.RegisteredUser.AccessKey(),
	}, nil
}
