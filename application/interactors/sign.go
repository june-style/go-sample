package interactors

import (
	"context"

	"github.com/june-style/go-sample/application/usecases"
	"github.com/june-style/go-sample/domain/dconfig"
	"github.com/june-style/go-sample/domain/derrors"
	"github.com/june-style/go-sample/domain/entities"
)

func NewSign(cfg *dconfig.Config, repository *entities.Repository) usecases.Sign {
	return &signImpl{
		config:       cfg,
		repositories: repository,
	}
}

type signImpl struct {
	config       *dconfig.Config
	repositories *entities.Repository
}

func (u *signImpl) In(ctx context.Context, input usecases.SignInInputData) (usecases.SignInOutputData, error) {
	userProfile, err := u.repositories.UserProfile.Find(ctx, input.UserID)
	if err != nil {
		return usecases.SignInOutputData{}, derrors.Wrap(err)
	}

	userSession, err := entities.CreateUserSession(userProfile.UserID(), u.config.Sys.SecretSalt)
	if err != nil {
		return usecases.SignInOutputData{}, derrors.Wrap(err)
	}
	if err := u.repositories.UserSession.Create(ctx, userSession); err != nil {
		return usecases.SignInOutputData{}, derrors.Wrap(err)
	}

	return usecases.SignInOutputData{
		UserSession: userSession,
	}, nil
}

func (u *signImpl) Up(ctx context.Context, input usecases.SignUpInputData) (usecases.SignUpOutputData, error) {
	var output usecases.SignUpOutputData

	if err := u.repositories.DynamoDB.TxWrite(ctx, func(ctx context.Context) error {
		registeredUser := entities.CreateRegisteredUser()
		if err := u.repositories.RegisteredUser.Create(ctx, registeredUser); err != nil {
			return derrors.Wrap(err)
		}
		userProfile := entities.NewUserProfile(registeredUser.UserID(), input.Sign, registeredUser.CreatedAt())
		if err := u.repositories.UserProfile.Create(ctx, userProfile); err != nil {
			return derrors.Wrap(err)
		}
		output.RegisteredUser = registeredUser
		return nil
	}); err != nil {
		return usecases.SignUpOutputData{}, derrors.Wrap(err)
	}

	return output, nil
}
