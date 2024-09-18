package derrors

import (
	"errors"
	"fmt"
	"net/http"
)

func Wrap(err error) error {
	return &wrapedError{
		err:   err,
		frame: Caller(1),
	}
}

func Wrapf(err error, format string, anys ...any) error {
	return &wrapedError{
		err:   err,
		msg:   fmt.Sprintf(format, anys...),
		frame: Caller(1),
	}
}

func New(message string) error {
	return &basicError{
		err:  errors.New(message),
		code: http.StatusInternalServerError,
		kind: ErrKindUnknown,
	}
}

func NewInvalidArgument(message string) error {
	return &basicError{
		err:  errors.New(message),
		code: http.StatusBadRequest,
		kind: ErrKindInvalidArgument,
	}
}

func NewNotFound(message string) error {
	return &basicError{
		err:  errors.New(message),
		code: http.StatusNotFound,
		kind: ErrKindNotFound,
	}
}

func NewPermissionDenied(message string) error {
	return &basicError{
		err:  errors.New(message),
		code: http.StatusForbidden,
		kind: ErrKindPermissionDenied,
	}
}

func NewInternal(message string) error {
	return &basicError{
		err:  errors.New(message),
		code: http.StatusInternalServerError,
		kind: ErrKindInternal,
	}
}

func NewUnauthenticated(message string) error {
	return &basicError{
		err:  errors.New(message),
		code: http.StatusUnauthorized,
		kind: ErrKindUnauthenticated,
	}
}
