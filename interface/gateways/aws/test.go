package aws

import (
	"time"

	resty "github.com/go-resty/resty/v2"
	"github.com/june-style/go-sample/domain/derrors"
)

const (
	LocalstackHealthCheckURL = "http://localhost:4566/health"

	localstackHealthCheckRetryCount = 30
	localstackHealthCheckWaitTime   = 1 * time.Second
)

var localstackHealthCheckDone bool

func CheckLocalstackServer() error {
	if localstackHealthCheckDone {
		return nil
	}

	c := resty.New()
	_, err := c.SetRetryCount(localstackHealthCheckRetryCount).
		SetRetryWaitTime(localstackHealthCheckWaitTime).
		AddRetryCondition(func(r *resty.Response, err error) bool {
			return r.StatusCode() != 200
		}).
		R().
		Get(LocalstackHealthCheckURL)
	if err != nil {
		return derrors.Wrap(err)
	}

	localstackHealthCheckDone = true
	return nil
}
