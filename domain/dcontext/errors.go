package dcontext

import "github.com/june-style/go-sample/domain/derrors"

var (
	ErrFailedToGetMetadata = derrors.NewInternal("failed to get metadata from incoming context")
	ErrFailedToGetTime     = derrors.NewInternal("failed to get time from incoming context")
)
