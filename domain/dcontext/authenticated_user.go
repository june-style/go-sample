package dcontext

import (
	"context"

	"github.com/june-style/go-sample/domain/entities"
)

type authenticatedUserContextKey struct{}

func GetAuthenticatedUserID(ctx context.Context) string {
	entity, ok := ctx.Value(authenticatedUserContextKey{}).(*entities.RegisteredUser)
	if !ok {
		return ""
	}
	return entity.UserID()
}

func GetAuthenticatedUser(ctx context.Context) *entities.RegisteredUser {
	entity, ok := ctx.Value(authenticatedUserContextKey{}).(*entities.RegisteredUser)
	if !ok {
		return nil
	}
	return entity
}

func SetAuthenticatedUser(ctx context.Context, entity *entities.RegisteredUser) context.Context {
	return context.WithValue(ctx, authenticatedUserContextKey{}, entity)
}
