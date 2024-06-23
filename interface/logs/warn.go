package logs

import (
	"github.com/june-style/go-sample/domain/derrors"
	"github.com/rs/zerolog"
)

func Warn() *warn {
	return &warn{
		loggers.warn.Warn(),
	}
}

type warn struct {
	*zerolog.Event
}

func (w *warn) SetRequestID(rid string) *warn {
	w.Str("rid", rid)
	return w
}

func (w *warn) SetError(err error) *warn {
	w.Err(err)
	if set := derrors.Stacktrace(err); len(set) > 0 {
		w.Interface("stacktrace", set)
	}
	return w
}
