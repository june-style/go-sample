package services_test

import (
	"context"
	"testing"
	"time"

	"github.com/june-style/go-sample/domain/dconfig"
	"github.com/june-style/go-sample/domain/dcontext"
	"github.com/june-style/go-sample/domain/entities"
	"github.com/june-style/go-sample/domain/services"
	"github.com/stretchr/testify/assert"
)

func Test_JWTer_Create(t *testing.T) {
	var testConfig = &dconfig.Config{
		App: dconfig.App{
			Key:                   "fake_application_key",
			SessionExpirationTime: 3600,
			HMACSecret:            "fake_hmac_secret",
		},
	}
	var testRegisteredUser = entities.NewRegisteredUser(
		entities.GenUUID(),
		entities.GenXID(),
		time.Date(2024, 12, 25, 0, 0, 0, 1, time.Local),
	)
	var testDate = time.Date(2024, 12, 25, 0, 0, 0, 1, time.Local)
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success with test data",
			args: args{
				ctx: func() context.Context {
					ctx := dcontext.SetAuthenticatedUser(context.TODO(), testRegisteredUser)
					return dcontext.SetTime(ctx, testDate)
				}(),
			},
			wantErr: assert.NoError,
		},
		{
			name: "failure with empty time",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: func(t assert.TestingT, err error, msgAndArgs ...any) bool {
				return assert.ErrorIs(t, err, dcontext.ErrFailedToGetTime, msgAndArgs...)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := services.NewJWTer(testConfig)
			_, err := j.Create(tt.args.ctx)
			if !tt.wantErr(t, err) {
				return
			}
		})
	}
}

func Test_JWTer_Verify(t *testing.T) {
	var testConfig = &dconfig.Config{
		App: dconfig.App{
			Key:                   "fake_application_key",
			SessionExpirationTime: 3600,
			HMACSecret:            "fake_hmac_secret",
		},
	}
	var testRegisteredUser = entities.NewRegisteredUser(
		entities.GenUUID(),
		entities.GenXID(),
		time.Date(2024, 12, 25, 0, 0, 0, 1, time.Local),
	)
	var testDate = time.Date(2024, 12, 25, 0, 0, 0, 1, time.Local)
	type args struct {
		ctx   context.Context
		token func() string
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success with test data",
			args: args{
				ctx: func() context.Context {
					ctx := dcontext.SetAuthenticatedUser(context.TODO(), testRegisteredUser)
					return dcontext.SetTime(ctx, testDate)
				}(),
				token: func() string {
					svc := services.NewJWTer(testConfig)
					ctx := dcontext.SetAuthenticatedUser(context.TODO(), testRegisteredUser)
					ctx = dcontext.SetTime(ctx, testDate)
					t, _ := svc.Create(ctx)
					return t
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "failure with token has invalid issuer",
			args: args{
				ctx: func() context.Context {
					ctx := dcontext.SetAuthenticatedUser(context.TODO(), testRegisteredUser)
					return dcontext.SetTime(ctx, testDate)
				}(),
				token: func() string {
					svc := services.NewJWTer(&dconfig.Config{
						App: dconfig.App{
							Key:                   "other_fake_application_key",
							SessionExpirationTime: testConfig.App.SessionExpirationTime,
							HMACSecret:            testConfig.App.HMACSecret,
						},
					})
					ctx := dcontext.SetAuthenticatedUser(context.TODO(), testRegisteredUser)
					ctx = dcontext.SetTime(ctx, testDate)
					t, _ := svc.Create(ctx)
					return t
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "failure with token is expired",
			args: args{
				ctx: func() context.Context {
					ctx := dcontext.SetAuthenticatedUser(context.TODO(), testRegisteredUser)
					return dcontext.SetTime(ctx, testDate)
				}(),
				token: func() string {
					svc := services.NewJWTer(&dconfig.Config{
						App: dconfig.App{
							Key:                   testConfig.App.Key,
							SessionExpirationTime: 0,
							HMACSecret:            testConfig.App.HMACSecret,
						},
					})
					ctx := dcontext.SetAuthenticatedUser(context.TODO(), testRegisteredUser)
					ctx = dcontext.SetTime(ctx, testDate)
					t, _ := svc.Create(ctx)
					return t
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "failure with signature is invalid",
			args: args{
				ctx: func() context.Context {
					ctx := dcontext.SetAuthenticatedUser(context.TODO(), testRegisteredUser)
					return dcontext.SetTime(ctx, testDate)
				}(),
				token: func() string {
					svc := services.NewJWTer(&dconfig.Config{
						App: dconfig.App{
							Key:                   testConfig.App.Key,
							SessionExpirationTime: testConfig.App.SessionExpirationTime,
							HMACSecret:            "other_fake_hmac_secret",
						},
					})
					ctx := dcontext.SetAuthenticatedUser(context.TODO(), testRegisteredUser)
					ctx = dcontext.SetTime(ctx, testDate)
					t, _ := svc.Create(ctx)
					return t
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "failure with token has invalid audience",
			args: args{
				ctx: func() context.Context {
					ctx := dcontext.SetAuthenticatedUser(context.TODO(), testRegisteredUser)
					return dcontext.SetTime(ctx, testDate)
				}(),
				token: func() string {
					svc := services.NewJWTer(testConfig)
					ctx := dcontext.SetAuthenticatedUser(context.TODO(), entities.NewRegisteredUser(
						testRegisteredUser.AccessKey(),
						"fake_xid",
						testRegisteredUser.CreatedAt(),
					))
					ctx = dcontext.SetTime(ctx, testDate)
					t, _ := svc.Create(ctx)
					return t
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "failure with empty time",
			args: args{
				ctx: func() context.Context {
					ctx := context.TODO()
					return dcontext.SetAuthenticatedUser(ctx, testRegisteredUser)
				}(),
				token: func() string {
					svc := services.NewJWTer(testConfig)
					ctx := dcontext.SetAuthenticatedUser(context.TODO(), testRegisteredUser)
					ctx = dcontext.SetTime(ctx, testDate)
					t, _ := svc.Create(ctx)
					return t
				},
			},
			wantErr: func(t assert.TestingT, err error, msgAndArgs ...any) bool {
				return assert.ErrorIs(t, err, dcontext.ErrFailedToGetTime, msgAndArgs...)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := services.NewJWTer(testConfig)
			err := j.Verify(tt.args.ctx, tt.args.token())
			if !tt.wantErr(t, err) {
				return
			}
		})
	}
}
