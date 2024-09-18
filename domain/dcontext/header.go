package dcontext

import (
	"context"

	"github.com/june-style/go-sample/domain/derrors"
)

const (
	HeaderApplicationKey Header = "x-application-key"
	HeaderAccessKey      Header = "x-access-key"
	HeaderSessionID      Header = "x-session-id"
	HeaderRequestID      Header = "x-request-id"
	HeaderUserAgent      Header = "user-agent"
)

type Header string

func (h Header) String() string {
	return string(h)
}

func GetHeader(ctx context.Context, header Header) (string, error) {
	md, err := GetMetadata(ctx)
	if err != nil {
		return "", derrors.Wrap(err)
	}
	if h := md.Get(header.String()); len(h) > 0 {
		return h[0], nil
	}
	return "", nil
}

var xHeaders = []Header{
	HeaderApplicationKey,
	HeaderAccessKey,
	HeaderSessionID,
	HeaderRequestID,
}

func GetHeaders(ctx context.Context) (map[string]string, error) {
	headers := make(map[string]string, 0)
	md, err := GetMetadata(ctx)
	if err != nil {
		return nil, derrors.Wrap(err)
	}
	for _, header := range xHeaders {
		if h := md.Get(header.String()); len(h) > 0 {
			headers[header.String()] = h[0]
		}
	}
	return headers, nil
}

func GetUserAgent(ctx context.Context) (string, error) {
	ua, err := GetHeader(ctx, HeaderUserAgent)
	if err != nil {
		return "", err
	}
	return ua, nil
}
