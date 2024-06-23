package repositories_test

import (
	"context"
	"testing"
	"time"

	"github.com/june-style/go-sample/domain/entities"
	"github.com/june-style/go-sample/interface/repositories"
	"github.com/stretchr/testify/assert"
)

func Test_registeredUserImpl_Find(t *testing.T) {
	var testRegisteredUser = entities.NewRegisteredUser(
		entities.GenUUID(),
		entities.GenXID(),
		time.Date(2024, 12, 25, 0, 0, 0, 1, time.Local),
	)
	type args struct {
		ctx       context.Context
		accessKey string
	}
	tests := []struct {
		name    string
		args    args
		want    *entities.RegisteredUser
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				ctx:       context.TODO(),
				accessKey: testRegisteredUser.AccessKey(),
			},
			want:    testRegisteredUser,
			wantErr: assert.NoError,
		},
		{
			name: "failure",
			args: args{
				ctx:       context.TODO(),
				accessKey: "",
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := repositories.NewRegisteredUserRepository(cfg, awsClient)
			if err != nil {
				assert.Fail(t, err.Error())
			}
			r.Create(tt.args.ctx, testRegisteredUser)
			t.Cleanup(func() { r.Delete(tt.args.ctx, testRegisteredUser) })

			got, err := r.Find(tt.args.ctx, tt.args.accessKey)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_registeredUserImpl_Create(t *testing.T) {
	var testRegisteredUser = entities.NewRegisteredUser(
		entities.GenUUID(),
		entities.GenXID(),
		time.Date(2024, 12, 25, 0, 0, 0, 1, time.Local),
	)
	type args struct {
		ctx            context.Context
		registeredUser *entities.RegisteredUser
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				ctx:            context.TODO(),
				registeredUser: testRegisteredUser,
			},
			wantErr: assert.NoError,
		},
		{
			name: "failure",
			args: args{
				ctx:            context.TODO(),
				registeredUser: nil,
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := repositories.NewRegisteredUserRepository(cfg, awsClient)
			if err != nil {
				assert.Fail(t, err.Error())
			}
			r.Create(tt.args.ctx, testRegisteredUser)
			t.Cleanup(func() { r.Delete(tt.args.ctx, testRegisteredUser) })

			tt.wantErr(t, r.Create(tt.args.ctx, tt.args.registeredUser))
		})
	}
}

func Test_registeredUserImpl_Delete(t *testing.T) {
	var testRegisteredUser = entities.NewRegisteredUser(
		entities.GenUUID(),
		entities.GenXID(),
		time.Date(2024, 12, 25, 0, 0, 0, 1, time.Local),
	)
	type args struct {
		ctx            context.Context
		registeredUser *entities.RegisteredUser
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				ctx:            context.TODO(),
				registeredUser: testRegisteredUser,
			},
			wantErr: assert.NoError,
		},
		{
			name: "failure",
			args: args{
				ctx:            context.TODO(),
				registeredUser: nil,
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := repositories.NewRegisteredUserRepository(cfg, awsClient)
			if err != nil {
				assert.Fail(t, err.Error())
			}
			r.Create(tt.args.ctx, testRegisteredUser)
			t.Cleanup(func() { r.Delete(tt.args.ctx, testRegisteredUser) })

			tt.wantErr(t, r.Delete(tt.args.ctx, tt.args.registeredUser))
		})
	}
}
