package redis

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/june-style/go-sample/domain/derrors"
	"github.com/redis/go-redis/v9"
)

type baseQuery interface {
	Del(ctx context.Context, id string) error
	Exists(ctx context.Context, id string) error
	Expire(ctx context.Context, id string, expir time.Duration) error
	ExpireAt(ctx context.Context, id string, time time.Time) error
}

type baseQueryImpl struct {
	db    redis.UniversalClient
	key   string
	expir time.Duration
}

func (b *baseQueryImpl) Del(ctx context.Context, id string) error {
	err := b.db.Del(ctx, b.getKey(id)).Err()
	if err != nil {
		return derrors.Wrap(err)
	}
	return nil
}

func (b *baseQueryImpl) Exists(ctx context.Context, id string) error {
	err := b.db.Exists(ctx, b.getKey(id)).Err()
	if err != nil {
		return derrors.Wrap(err)
	}
	return nil
}

func (b *baseQueryImpl) Expire(ctx context.Context, id string, expir time.Duration) error {
	err := b.db.Expire(ctx, b.getKey(id), expir).Err()
	if err != nil {
		return derrors.Wrap(err)
	}
	return nil
}

func (b *baseQueryImpl) ExpireAt(ctx context.Context, id string, time time.Time) error {
	err := b.db.ExpireAt(ctx, b.getKey(id), time).Err()
	if err != nil {
		return derrors.Wrap(err)
	}
	return nil
}

func (b *baseQueryImpl) getKey(suffixes ...string) string {
	if len(suffixes) == 0 {
		return b.key
	}
	return fmt.Sprintf("%s:%s", b.key, strings.Join(suffixes, ":"))
}
