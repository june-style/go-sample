package interactors_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/june-style/go-sample/application/interactors"
	"github.com/june-style/go-sample/application/usecases"
	"github.com/june-style/go-sample/domain/derrors"
	"github.com/june-style/go-sample/domain/entities"
	entities_mock "github.com/june-style/go-sample/domain/entities/mock"
	"github.com/stretchr/testify/assert"
)

func Test_homeImpl_Get(t *testing.T) {
	var testUserProfile = entities.NewUserProfile(
		entities.GenXID(),
		"dummy-user-name",
		time.Date(2024, 12, 25, 0, 0, 0, 1, time.Local),
	)
	type fields struct {
		newUserProfileRepository func(ctrl *gomock.Controller) entities.UserProfileRepository
	}
	type args struct {
		ctx   context.Context
		input usecases.HomeGetInputData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    usecases.HomeGetOutputData
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			fields: fields{
				newUserProfileRepository: func(ctrl *gomock.Controller) entities.UserProfileRepository {
					m := entities_mock.NewMockUserProfileRepository(ctrl)
					m.EXPECT().Find(gomock.Any(), testUserProfile.UserID()).Return(testUserProfile, nil)
					return m
				},
			},
			args: args{
				ctx: context.TODO(),
				input: usecases.HomeGetInputData{
					UserID: testUserProfile.UserID(),
				},
			},
			want: usecases.HomeGetOutputData{
				UserProfile: testUserProfile,
			},
			wantErr: assert.NoError,
		},
		{
			name: "failure",
			fields: fields{
				newUserProfileRepository: func(ctrl *gomock.Controller) entities.UserProfileRepository {
					m := entities_mock.NewMockUserProfileRepository(ctrl)
					m.EXPECT().Find(gomock.Any(), testUserProfile.UserID()).Return(nil, derrors.New("unknown error"))
					return m
				},
			},
			args: args{
				ctx: context.TODO(),
				input: usecases.HomeGetInputData{
					UserID: testUserProfile.UserID(),
				},
			},
			want: usecases.HomeGetOutputData{
				UserProfile: nil,
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			t.Cleanup(func() { ctrl.Finish() })

			u := interactors.NewHome(&entities.Repository{
				UserProfile: tt.fields.newUserProfileRepository(ctrl),
			})
			got, err := u.Get(tt.args.ctx, tt.args.input)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
