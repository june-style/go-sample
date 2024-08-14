package interceptors

import (
	"context"

	"github.com/june-style/go-sample/domain/derrors"
	"github.com/june-style/go-sample/framework/protocol/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func ErrorStatus() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		res, err := handler(ctx, req)

		if err != nil {
			kind := derrors.Kind(err)
			st := status.New(kind.Code, kind.Msg)
			st, sErr := st.WithDetails(
				ErrorDetailPresenter(err),
			)
			if sErr != nil {
				return nil, sErr
			}
			return nil, st.Err()
		}

		return res, nil
	}
}

func ErrorDetailPresenter(err error) *pb.ErrorDetail {
	return &pb.ErrorDetail{
		Code:       int32(derrors.Code(err)),
		Msg:        err.Error(),
		Stacktrace: derrors.Stacktrace(err),
	}
}
