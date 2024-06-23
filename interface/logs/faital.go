package logs

import (
	"github.com/june-style/go-sample/domain/derrors"
	"github.com/rs/zerolog"
)

func Fatal() *fatal {
	return &fatal{
		loggers.err.Fatal(),
	}
}

type fatal struct {
	*zerolog.Event
}

func (f *fatal) SetRequestID(rid string) *fatal {
	f.Str("rid", rid)
	return f
}

func (f *fatal) SetError(err error) *fatal {
	f.Err(err)
	if set := derrors.Stacktrace(err); len(set) > 0 {
		f.Interface("stacktrace", set)
	}
	return f
}
