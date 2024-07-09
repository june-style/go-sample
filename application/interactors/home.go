package interactors

import (
	"context"

	"github.com/june-style/go-sample/application/usecases"
	"github.com/june-style/go-sample/domain/derrors"
	"github.com/june-style/go-sample/domain/entities"
)

func NewHome(repository *entities.Repository) usecases.Home {
	return &homeImpl{
		repositories: repository,
	}
}

type homeImpl struct {
	repositories *entities.Repository
}

func (u *homeImpl) Get(ctx context.Context, input usecases.HomeGetInputData) (usecases.HomeGetOutputData, error) {
	var output usecases.HomeGetOutputData

	userProfile, err := u.repositories.UserProfile.Find(ctx, input.UserID)
	if err != nil {
		return output, derrors.Wrap(err)
	}
	output.UserProfile = userProfile

	return output, nil
}
