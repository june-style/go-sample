package dcontext

import (
	"context"
	"time"

	"github.com/june-style/go-sample/domain/derrors"
)

type timeContextKey struct{}

func GetTime(ctx context.Context) (time.Time, error) {
	if time, ok := ctx.Value(timeContextKey{}).(time.Time); ok {
		return time, nil
	}
	return time.Time{}, derrors.Wrap(ErrFailedToGetTime)
}

func SetTime(ctx context.Context, time time.Time) context.Context {
	return context.WithValue(ctx, timeContextKey{}, time)
}
