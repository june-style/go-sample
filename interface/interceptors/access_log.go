package interceptors

import (
	"context"

	"github.com/june-style/go-sample/domain/dcontext"
	"github.com/june-style/go-sample/domain/derrors"
	"github.com/june-style/go-sample/interface/logs"
	"google.golang.org/grpc"
)

func AccessLog() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		ua, err := dcontext.GetUserAgent(ctx)
		if err != nil {
			return nil, derrors.Wrap(err)
		}

		rid, err := dcontext.GetHeader(ctx, dcontext.HeaderRequestID)
		if err != nil {
			return nil, derrors.Wrap(err)
		}

		headers, err := dcontext.GetHeaders(ctx)
		if err != nil {
			return nil, derrors.Wrap(err)
		}

		logs.Access().
			SetRequestID(rid).
			SetMethod(info.FullMethod).
			SetRequest(req).
			SetHeader(headers).
			SetUserAgent(ua).
			Msg("Access")

		return handler(ctx, req)
	}
}
