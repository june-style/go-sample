package repositories

import (
	"context"
	"time"

	"github.com/june-style/go-sample/domain/derrors"
	"github.com/june-style/go-sample/domain/entities"
	"github.com/june-style/go-sample/interface/gateways/redis"
)

type UserSession struct {
	UserID    string `json:"uid"`
	SessionID string `json:"sid"`
}

type UserSessionSet []*UserSession

func NewUserSessionRepository(client *redis.Client) (entities.UserSessionRepository, error) {
	q, err := client.NewQueryString(redis.QueryInputData{
		KeyName:    "user-session",
		DBIndex:    redis.DB_00,
		ExpireTime: time.Hour * 24,
	})
	if err != nil {
		return nil, err
	}
	return &userSessionImpl{
		query: q,
	}, nil
}

type userSessionImpl struct {
	query redis.StringQuery
}

func (u *userSessionImpl) Find(ctx context.Context, userID string) (*entities.UserSession, error) {
	if userID == "" {
		return nil, derrors.Wrap(derrors.NewInvalidArgument("invalid parameter for find user session"))
	}
	sessionID, err := u.query.Get(ctx, userID)
	if err != nil {
		return nil, derrors.Wrap(err)
	}
	return entities.NewUserSession(
		userID,
		sessionID,
	), nil
}

func (u *userSessionImpl) Create(ctx context.Context, userSession *entities.UserSession) error {
	if userSession.UserID() == "" || userSession.SessionID() == "" {
		return derrors.Wrap(derrors.NewInvalidArgument("invalid parameters for create user session"))
	}
	if err := u.query.Set(ctx, userSession.UserID(), userSession.SessionID()); err != nil {
		return derrors.Wrap(err)
	}
	return nil
}

func (u *userSessionImpl) Delete(ctx context.Context, userSession *entities.UserSession) error {
	if userSession.UserID() == "" || userSession.SessionID() == "" {
		return derrors.Wrap(derrors.NewInvalidArgument("invalid parameters for delete user session"))
	}
	if err := u.query.Del(ctx, userSession.UserID()); err != nil {
		return derrors.Wrap(err)
	}
	return nil
}
