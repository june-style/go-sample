package repositories_test

import (
	"context"
	"testing"

	"github.com/june-style/go-sample/domain/entities"
	"github.com/june-style/go-sample/interface/repositories"
	"github.com/stretchr/testify/assert"
)

func Test_userSessionImpl_Find(t *testing.T) {
	var testUserSession, _ = entities.CreateUserSession(
		entities.GenXID(),
		cfg.Sys.SecretSalt,
	)
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		args    args
		want    *entities.UserSession
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				ctx:    context.TODO(),
				userID: testUserSession.UserID(),
			},
			want:    testUserSession,
			wantErr: assert.NoError,
		},
		{
			name: "failure",
			args: args{
				ctx:    context.TODO(),
				userID: "",
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := repositories.NewUserSessionRepository(redisClient)
			if err != nil {
				assert.Fail(t, err.Error())
			}
			r.Create(tt.args.ctx, testUserSession)
			t.Cleanup(func() { r.Delete(tt.args.ctx, testUserSession) })

			got, err := r.Find(tt.args.ctx, tt.args.userID)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_userSessionImpl_Create(t *testing.T) {
	var testUserSession, _ = entities.CreateUserSession(
		entities.GenXID(),
		cfg.Sys.SecretSalt,
	)
	type args struct {
		ctx         context.Context
		userSession *entities.UserSession
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				ctx:         context.TODO(),
				userSession: testUserSession,
			},
			wantErr: assert.NoError,
		},
		{
			name: "failure",
			args: args{
				ctx:         context.TODO(),
				userSession: nil,
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := repositories.NewUserSessionRepository(redisClient)
			if err != nil {
				assert.Fail(t, err.Error())
			}
			r.Create(tt.args.ctx, testUserSession)
			t.Cleanup(func() { r.Delete(tt.args.ctx, testUserSession) })

			tt.wantErr(t, r.Create(tt.args.ctx, tt.args.userSession))
		})
	}
}

func Test_userSessionImpl_Delete(t *testing.T) {
	var testUserSession, _ = entities.CreateUserSession(
		entities.GenXID(),
		cfg.Sys.SecretSalt,
	)
	type args struct {
		ctx         context.Context
		userSession *entities.UserSession
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				ctx:         context.TODO(),
				userSession: testUserSession,
			},
			wantErr: assert.NoError,
		},
		{
			name: "failure",
			args: args{
				ctx:         context.TODO(),
				userSession: nil,
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := repositories.NewUserSessionRepository(redisClient)
			if err != nil {
				assert.Fail(t, err.Error())
			}
			r.Create(tt.args.ctx, testUserSession)
			t.Cleanup(func() { r.Delete(tt.args.ctx, testUserSession) })

			tt.wantErr(t, r.Delete(tt.args.ctx, tt.args.userSession))
		})
	}
}
