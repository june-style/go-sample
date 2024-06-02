package entities

import "github.com/june-style/go-sample/domain/derrors"

var (
	ErrParameterDoesNotExists = derrors.NewInvalidArgument("parameter does not exist")
	ErrParameterAlreadyExists = derrors.NewInvalidArgument("parameter already exists")
)

var (
	ErrExceedsMinimumValue = derrors.NewInvalidArgument("exceeds minimum value")
	ErrExceedsMaximumValue = derrors.NewInvalidArgument("exceeds maximum value")
)
