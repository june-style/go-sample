package redis

import (
	"context"

	"github.com/june-style/go-sample/domain/derrors"
)

func (c *Client) NewQueryString(q QueryInputData) (StringQuery, error) {
	db, err := c.GetDB(q.GetDBIndex())
	if err != nil {
		return nil, err
	}
	return &stringQueryImpl{baseQueryImpl{
		db:    db,
		key:   c.GetKeyFullName(q.GetKeyName()),
		expir: q.GetExpireTime(),
	}}, nil
}

type StringQuery interface {
	baseQuery
	Set(ctx context.Context, id string, value any) error
	Get(ctx context.Context, id string) (string, error)
}

type stringQueryImpl struct {
	baseQueryImpl
}

func (s *stringQueryImpl) Set(ctx context.Context, id string, value any) error {
	err := s.db.Set(ctx, s.getKey(id), value, s.expir).Err()
	if err != nil {
		return derrors.Wrap(err)
	}
	return nil
}

func (s *stringQueryImpl) Get(ctx context.Context, id string) (string, error) {
	res, err := s.db.Get(ctx, s.getKey(id)).Result()
	if err != nil {
		return "", derrors.Wrap(err)
	}
	return res, nil
}
