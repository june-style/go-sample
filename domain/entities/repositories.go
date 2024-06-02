package entities

import "context"

type Repository struct {
	// tx
	DynamoDB DynamoDB
	// users
	RegisteredUser RegisteredUserRepository
	UserProfile    UserProfileRepository
	UserSession    UserSessionRepository
}

//go:generate mockgen -source=${GOFILE} -destination=./mock/${GOFILE} -package=${GOPACKAGE}_mock
type RegisteredUserRepository interface {
	Find(ctx context.Context, accessKey string) (*RegisteredUser, error)
	Create(ctx context.Context, RegisteredUser *RegisteredUser) error
	Delete(ctx context.Context, RegisteredUser *RegisteredUser) error
}

//go:generate mockgen -source=${GOFILE} -destination=./mock/${GOFILE} -package=${GOPACKAGE}_mock
type UserProfileRepository interface {
	Find(ctx context.Context, userID string) (*UserProfile, error)
	Create(ctx context.Context, userProfile *UserProfile) error
	Delete(ctx context.Context, userProfile *UserProfile) error
}

//go:generate mockgen -source=${GOFILE} -destination=./mock/${GOFILE} -package=${GOPACKAGE}_mock
type UserSessionRepository interface {
	Find(ctx context.Context, userID string) (*UserSession, error)
	Create(ctx context.Context, userSession *UserSession) error
	Delete(ctx context.Context, userSession *UserSession) error
}
