package usecases

import (
	"context"

	"github.com/june-style/go-sample/domain/entities"
)

//go:generate mockgen -source=${GOFILE} -destination=./mock/${GOFILE} -package=${GOPACKAGE}_mock
type Sign interface {
	In(ctx context.Context, input SignInInputData) (SignInOutputData, error)
	Up(ctx context.Context, input SignUpInputData) (SignUpOutputData, error)
}

type SignInInputData struct {
	UserID string
}

type SignInOutputData struct {
	UserSession *entities.UserSession
}

type SignUpInputData struct {
	Sign string
}

type SignUpOutputData struct {
	RegisteredUser *entities.RegisteredUser
}
