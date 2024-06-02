package dconfig_test

import (
	"testing"

	"github.com/june-style/go-sample/domain/dconfig"
	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	var testConfig = &dconfig.Config{
		App: dconfig.App{
			Name:                  "june-style-go-sample",
			Key:                   "fake_application_key",
			SessionExpirationTime: 3600,
			HMACSecret:            "fake_hmac_secret",
		},
		Aws: dconfig.Aws{
			AccessKeyID:     "fake_access_key",
			SecretAccessKey: "fake_secret_access_key",
			Region:          "ap-northeast-1",
			Endpoint:        "http://localstack:4566",
			DynamoDBTable:   "data",
			DebugMode:       false,
		},
		Grpc: dconfig.Grpc{
			Port: 9000,
		},
		Redis: dconfig.Redis{
			Server:   "redis:6379",
			DBNumber: 16,
		},
		Sys: dconfig.Sys{
			Env:        "localhost",
			TZ:         "Asia/Tokyo",
			SecretSalt: "fake_secret_salt",
		},
	}
	type args struct {
		dotenv string
	}
	tests := []struct {
		name    string
		args    args
		want    *dconfig.Config
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				dotenv: "../../.env.template",
			},
			want:    testConfig,
			wantErr: assert.NoError,
		},
		{
			name: "failure",
			args: args{
				dotenv: "",
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dconfig.New(tt.args.dotenv)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
