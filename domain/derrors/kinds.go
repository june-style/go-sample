package derrors

import "google.golang.org/grpc/codes"

type ErrKind struct {
	Code codes.Code
	Msg  string
}

var (
	ErrKindUnknown = ErrKind{
		Code: codes.Unknown,
		Msg:  "HTTP: 500 Internal Server Error.",
	}

	ErrKindInvalidArgument = ErrKind{
		Code: codes.InvalidArgument,
		Msg:  "HTTP: 400 Bad Request.",
	}

	ErrKindNotFound = ErrKind{
		Code: codes.NotFound,
		Msg:  "HTTP: 404 Not Found.",
	}

	ErrKindPermissionDenied = ErrKind{
		Code: codes.PermissionDenied,
		Msg:  "HTTP: 403 Forbidden.",
	}

	ErrKindInternal = ErrKind{
		Code: codes.Internal,
		Msg:  "HTTP: 500 Internal Server Error.",
	}

	ErrKindUnauthenticated = ErrKind{
		Code: codes.Unauthenticated,
		Msg:  "HTTP: 401 Unauthorized.",
	}
)
