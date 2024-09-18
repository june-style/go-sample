package logs

import (
	"github.com/june-style/go-sample/domain/derrors"
	"github.com/rs/zerolog"
)

func Panic() *pnc {
	return &pnc{
		loggers.err.Panic(),
	}
}

type pnc struct {
	*zerolog.Event
}

func (p *pnc) SetRequestID(rid string) *pnc {
	p.Str("rid", rid)
	return p
}

func (p *pnc) SetError(err error) *pnc {
	p.Err(err)
	if set := derrors.Stacktrace(err); len(set) > 0 {
		p.Interface("stacktrace", set)
	}
	return p
}
