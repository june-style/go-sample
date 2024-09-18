package logs

import (
	"github.com/june-style/go-sample/domain/derrors"
	"github.com/rs/zerolog"
)

func Error() *err {
	return &err{
		loggers.err.Error(),
	}
}

type err struct {
	*zerolog.Event
}

func (e *err) SetRequestID(rid string) *err {
	e.Str("rid", rid)
	return e
}

func (e *err) SetError(err error) *err {
	e.Err(err)
	if set := derrors.Stacktrace(err); len(set) > 0 {
		e.Interface("stacktrace", set)
	}
	return e
}
