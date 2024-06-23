package logs

import "github.com/rs/zerolog"

func Debug() *debug {
	return &debug{
		loggers.info.Debug(),
	}
}

type debug struct {
	*zerolog.Event
}

func (d *debug) SetRequestID(rid string) *debug {
	d.Str("rid", rid)
	return d
}

func (d *debug) SetResponse(res any) *debug {
	d.Interface("res", res)
	return d
}

func (d *debug) SetField(key string, value any) *debug {
	d.Interface(key, value)
	return d
}
