package logs

import "github.com/rs/zerolog"

func Info() *info {
	return &info{
		loggers.info.Info(),
	}
}

type info struct {
	*zerolog.Event
}

func (i *info) SetRequestID(rid string) *info {
	i.Str("rid", rid)
	return i
}

// Deprecated: Avoid using *zerolog.Event.Interface()
func (i *info) SetField(key string, value any) *info {
	i.Interface(key, value)
	return i
}
