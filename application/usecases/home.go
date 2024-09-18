package usecases

import (
	"context"

	"github.com/june-style/go-sample/domain/entities"
)

//go:generate mockgen -source=${GOFILE} -destination=./mock/${GOFILE} -package=${GOPACKAGE}_mock
type Home interface {
	Get(ctx context.Context, input HomeGetInputData) (HomeGetOutputData, error)
}

type HomeGetInputData struct {
	UserID string
}

type HomeGetOutputData struct {
	UserProfile *entities.UserProfile
}
