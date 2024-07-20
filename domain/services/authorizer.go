package services

import (
	"context"

	"github.com/june-style/go-sample/domain/dconfig"
	"github.com/june-style/go-sample/domain/dcontext"
	"github.com/june-style/go-sample/domain/derrors"
	"github.com/june-style/go-sample/domain/entities"
)

func NewAuthorizer(cfg *dconfig.Config, repository *entities.Repository) Authorizer {
	return &authorizerImpl{
		config:       cfg,
		repositories: repository,
	}
}

//go:generate mockgen -source=${GOFILE} -destination=./mock/${GOFILE} -package=${GOPACKAGE}_mock
type Authorizer interface {
	CreateAccessKey(ctx context.Context) (*entities.RegisteredUser, error)
	CreateSession(ctx context.Context, userID string) (*entities.UserSession, error)
	VerifyApplicationKey(ctx context.Context) error
	VerifyAccessKey(ctx context.Context) (context.Context, error)
	VerifySession(ctx context.Context) error
}

var (
	ErrIllegalApplicationKey = derrors.NewUnauthenticated("illegal application key")
	ErrIllegalAccessKey      = derrors.NewUnauthenticated("illegal access key")
	ErrIllegalSessionID      = derrors.NewUnauthenticated("illegal session ID")
)

type authorizerImpl struct {
	config       *dconfig.Config
	repositories *entities.Repository
}

func (s *authorizerImpl) CreateAccessKey(ctx context.Context) (*entities.RegisteredUser, error) {
	registeredUser := entities.CreateRegisteredUser()
	if err := s.repositories.RegisteredUser.Create(ctx, registeredUser); err != nil {
		return nil, derrors.Wrap(err)
	}
	return registeredUser, nil
}

func (s *authorizerImpl) CreateSession(ctx context.Context, userID string) (*entities.UserSession, error) {
	userSession, err := entities.CreateUserSession(userID, s.config.Sys.SecretSalt)
	if err != nil {
		return nil, derrors.Wrap(err)
	}
	if err := s.repositories.UserSession.Create(ctx, userSession); err != nil {
		return nil, derrors.Wrap(err)
	}
	return userSession, nil
}

func (s *authorizerImpl) VerifyApplicationKey(ctx context.Context) error {
	applicationKey, err := dcontext.GetHeader(ctx, dcontext.HeaderApplicationKey)
	if err != nil {
		return derrors.Wrap(err)
	}
	if s.config.App.Key != applicationKey {
		return derrors.Wrapf(ErrIllegalApplicationKey, "application-key is %s", applicationKey)
	}
	return nil
}

func (s *authorizerImpl) VerifyAccessKey(ctx context.Context) (context.Context, error) {
	accessKey, err := dcontext.GetHeader(ctx, dcontext.HeaderAccessKey)
	if err != nil {
		return nil, derrors.Wrap(err)
	}
	registeredUser, err := s.repositories.RegisteredUser.Find(ctx, accessKey)
	if err != nil {
		return nil, derrors.Wrap(err)
	}
	if registeredUser.AccessKey() != accessKey {
		return nil, derrors.Wrapf(ErrIllegalAccessKey, "access-key is %s", accessKey)
	}
	return dcontext.SetAuthenticatedUser(ctx, registeredUser), nil
}

func (s *authorizerImpl) VerifySession(ctx context.Context) error {
	userID := dcontext.GetAuthenticatedUserID(ctx)
	userSession, err := s.repositories.UserSession.Find(ctx, userID)
	if err != nil {
		return derrors.Wrap(err)
	}
	sessionID, err := dcontext.GetHeader(ctx, dcontext.HeaderSessionID)
	if err != nil {
		return derrors.Wrap(err)
	}
	if userSession.SessionID() != sessionID {
		return derrors.Wrapf(ErrIllegalSessionID, "session-id is %s", sessionID)
	}
	return nil
}
