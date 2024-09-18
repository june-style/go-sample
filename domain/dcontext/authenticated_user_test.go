package dcontext_test

import (
	"context"
	"testing"
	"time"

	"github.com/june-style/go-sample/domain/dcontext"
	"github.com/june-style/go-sample/domain/entities"
	"github.com/stretchr/testify/assert"
)

var testRegisteredUser = entities.NewRegisteredUser(
	entities.GenUUID(),
	entities.GenXID(),
	time.Date(2024, 12, 25, 0, 0, 0, 1, time.Local),
)

func Test_GetAuthenticatedUserID(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success with test data",
			args: args{
				ctx: dcontext.SetAuthenticatedUser(context.TODO(), testRegisteredUser),
			},
			want: testRegisteredUser.UserID(),
		},
		{
			name: "success with empty data",
			args: args{
				ctx: context.TODO(),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, dcontext.GetAuthenticatedUserID(tt.args.ctx))
		})
	}
}

func Test_GetAuthenticatedUser(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want *entities.RegisteredUser
	}{
		{
			name: "success with test data",
			args: args{
				ctx: dcontext.SetAuthenticatedUser(context.TODO(), testRegisteredUser),
			},
			want: testRegisteredUser,
		},
		{
			name: "success with empty data",
			args: args{
				ctx: context.TODO(),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, dcontext.GetAuthenticatedUser(tt.args.ctx))
		})
	}
}

func Test_SetAuthenticatedUser(t *testing.T) {
	type args struct {
		ctx    context.Context
		entity *entities.RegisteredUser
	}
	tests := []struct {
		name string
		args args
		want *entities.RegisteredUser
	}{
		{
			name: "success with test data",
			args: args{
				ctx:    context.TODO(),
				entity: testRegisteredUser,
			},
			want: testRegisteredUser,
		},
		{
			name: "success with empty data",
			args: args{
				ctx:    context.TODO(),
				entity: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				ctx := dcontext.SetAuthenticatedUser(tt.args.ctx, tt.args.entity)
				assert.Equal(t, tt.want, dcontext.GetAuthenticatedUser(ctx))
			})
		})
	}
}
