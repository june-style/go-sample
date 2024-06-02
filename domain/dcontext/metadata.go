package dcontext

import (
	"context"

	"github.com/june-style/go-sample/domain/derrors"
	"google.golang.org/grpc/metadata"
)

func GetMetadata(ctx context.Context) (metadata.MD, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, derrors.Wrap(ErrFailedToGetMetadata)
	}
	return md, nil
}

func SetMetadata(ctx context.Context, md metadata.MD) context.Context {
	return metadata.NewIncomingContext(ctx, md)
}
