package interceptors

import (
	"context"
	"fmt"

	"github.com/june-style/go-sample/domain/dcontext"
	"github.com/june-style/go-sample/domain/derrors"
	"github.com/june-style/go-sample/interface/logs"
	"google.golang.org/grpc"
)

func Recovery() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, err error) {
		panicked := true

		defer func() {
			if r := recover(); r != nil || panicked {
				if perr, ok := r.(error); ok {
					err = derrors.Wrap(perr)
				} else {
					err = derrors.Wrap(fmt.Errorf("%+v", r))
				}

				// access_log.go で既にコール済のためここでのエラーハンドリングはスキップする
				rid, _ := dcontext.GetHeader(ctx, dcontext.HeaderRequestID)

				logs.Error().
					SetRequestID(rid).
					SetError(err).
					Msg("Panic recovered")

				// TODO: panic用 後ほど綺麗に描く
				logs.Error().Interface("stacktrace", derrors.StacktraceAll()).Send()
			}
		}()

		res, err := handler(ctx, req)
		panicked = false
		return res, err
	}
}
