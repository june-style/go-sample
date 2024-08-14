package interceptors

import (
	"context"
	"time"

	"github.com/june-style/go-sample/domain/services"
	"google.golang.org/grpc"
)

const timeout = 30 * time.Second

func Initialization(timer services.Timer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		ctx = timer.SetNow(ctx)

		return handler(ctx, req)
	}
}
