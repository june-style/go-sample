package repositories_test

import (
	"context"
	"testing"
	"time"

	"github.com/june-style/go-sample/domain/entities"
	"github.com/june-style/go-sample/interface/repositories"
	"github.com/stretchr/testify/assert"
)

func Test_userProfileImpl_Find(t *testing.T) {
	var testUserProfile = entities.NewUserProfile(
		entities.GenXID(),
		"dummy-user-name",
		time.Date(2024, 12, 25, 0, 0, 0, 1, time.Local),
	)
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		args    args
		want    *entities.UserProfile
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				ctx:    context.TODO(),
				userID: testUserProfile.UserID(),
			},
			want:    testUserProfile,
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
			r, err := repositories.NewUserProfileRepository(cfg, awsClient)
			if err != nil {
				assert.Fail(t, err.Error())
			}
			r.Create(tt.args.ctx, testUserProfile)
			t.Cleanup(func() { r.Delete(tt.args.ctx, testUserProfile) })

			got, err := r.Find(tt.args.ctx, tt.args.userID)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_userProfileImpl_Create(t *testing.T) {
	var testUserProfile = entities.NewUserProfile(
		entities.GenXID(),
		"dummy-user-name",
		time.Date(2024, 12, 25, 0, 0, 0, 1, time.Local),
	)
	type args struct {
		ctx         context.Context
		userProfile *entities.UserProfile
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
				userProfile: testUserProfile,
			},
			wantErr: assert.NoError,
		},
		{
			name: "failure",
			args: args{
				ctx:         context.TODO(),
				userProfile: nil,
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := repositories.NewUserProfileRepository(cfg, awsClient)
			if err != nil {
				assert.Fail(t, err.Error())
			}
			t.Cleanup(func() { r.Delete(tt.args.ctx, testUserProfile) })

			tt.wantErr(t, r.Create(tt.args.ctx, tt.args.userProfile))
		})
	}
}

func Test_userProfileImpl_Delete(t *testing.T) {
	var testUserProfile = entities.NewUserProfile(
		entities.GenXID(),
		"dummy-user-name",
		time.Date(2024, 12, 25, 0, 0, 0, 1, time.Local),
	)
	type args struct {
		ctx         context.Context
		userProfile *entities.UserProfile
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
				userProfile: testUserProfile,
			},
			wantErr: assert.NoError,
		},
		{
			name: "failure",
			args: args{
				ctx:         context.TODO(),
				userProfile: nil,
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := repositories.NewUserProfileRepository(cfg, awsClient)
			if err != nil {
				assert.Fail(t, err.Error())
			}
			r.Create(tt.args.ctx, testUserProfile)
			t.Cleanup(func() { r.Delete(tt.args.ctx, testUserProfile) })

			tt.wantErr(t, r.Delete(tt.args.ctx, tt.args.userProfile))
		})
	}
}
