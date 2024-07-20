package controllers

import (
	"context"

	"github.com/june-style/go-sample/application/usecases"
	"github.com/june-style/go-sample/domain/dcontext"
	"github.com/june-style/go-sample/domain/derrors"
	"github.com/june-style/go-sample/framework/protocol/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewHome(usecase *usecases.UseCase) Home {
	return &homeImpl{
		usecases: usecase,
	}
}

type Home interface {
	pb.HomeServiceServer
}

type homeImpl struct {
	pb.UnimplementedHomeServiceServer
	usecases *usecases.UseCase
}

func (c *homeImpl) Get(ctx context.Context, req *pb.HomeGetRequest) (*pb.HomeGetResponse, error) {
	output, err := c.usecases.Home.Get(ctx, usecases.HomeGetInputData{
		UserID: dcontext.GetAuthenticatedUserID(ctx),
	})
	if err != nil {
		return nil, derrors.Wrap(err)
	}
	return &pb.HomeGetResponse{
		User: &pb.User{
			Id:        output.UserProfile.UserID(),
			Name:      output.UserProfile.Name(),
			CreatedAt: timestamppb.New(output.UserProfile.CreatedAt()),
		},
	}, nil
}
