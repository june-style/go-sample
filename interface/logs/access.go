package logs

import "github.com/rs/zerolog"

func Access() *access {
	return &access{
		loggers.access.Info(),
	}
}

type access struct {
	*zerolog.Event
}

func (a *access) SetRequestID(rid string) *access {
	a.Str("rid", rid)
	return a
}

func (a *access) SetMethod(method string) *access {
	a.Str("method", method)
	return a
}

func (a *access) SetUserID(uid string) *access {
	a.Str("uid", uid)
	return a
}

func (a *access) SetRequest(req any) *access {
	a.Interface("req", req)
	return a
}

func (a *access) SetHeader(header any) *access {
	a.Interface("header", header)
	return a
}

func (a *access) SetUserAgent(ua string) *access {
	a.Str("ua", ua)
	return a
}
