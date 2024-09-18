package interceptors

import (
	"context"

	"github.com/june-style/go-sample/domain/dcontext"
	"github.com/june-style/go-sample/interface/logs"
	"google.golang.org/grpc"
)

func ErrorLog() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		res, err := handler(ctx, req)

		if err != nil {
			// access_log.go で既にコール済のためここでのエラーハンドリングはスキップする
			rid, _ := dcontext.GetHeader(ctx, dcontext.HeaderRequestID)

			logs.Error().
				SetRequestID(rid).
				SetError(err).
				Msg("Error")
		}

		return res, err
	}
}
