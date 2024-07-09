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

func Test_signImpl_In(t *testing.T) {
	var testUserProfile = entities.NewUserProfile(
		entities.GenXID(),
		"dummy-user-name",
		time.Date(2024, 12, 25, 0, 0, 0, 1, time.Local),
	)
	type fields struct {
		newUserProfileRepository func(ctrl *gomock.Controller) entities.UserProfileRepository
		newUserSessionRepository func(ctrl *gomock.Controller) entities.UserSessionRepository
	}
	type args struct {
		ctx   context.Context
		input usecases.SignInInputData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    usecases.SignInOutputData
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
				newUserSessionRepository: func(ctrl *gomock.Controller) entities.UserSessionRepository {
					m := entities_mock.NewMockUserSessionRepository(ctrl)
					m.EXPECT().Create(gomock.Any(), gomock.AssignableToTypeOf(&entities.UserSession{})).Return(nil)
					return m
				},
			},
			args: args{
				ctx: context.TODO(),
				input: usecases.SignInInputData{
					UserID: testUserProfile.UserID(),
				},
			},
			want: usecases.SignInOutputData{
				UserSession: entities.NewUserSession(testUserProfile.UserID(), "dummy"),
			},
			wantErr: assert.NoError,
		},
		{
			name: "failure",
			fields: fields{
				newUserProfileRepository: func(ctrl *gomock.Controller) entities.UserProfileRepository {
					m := entities_mock.NewMockUserProfileRepository(ctrl)
					m.EXPECT().Find(gomock.Any(), testUserProfile.UserID()).Return(testUserProfile, nil)
					return m
				},
				newUserSessionRepository: func(ctrl *gomock.Controller) entities.UserSessionRepository {
					m := entities_mock.NewMockUserSessionRepository(ctrl)
					m.EXPECT().Create(gomock.Any(), gomock.AssignableToTypeOf(&entities.UserSession{})).Return(derrors.New("unknown error"))
					return m
				},
			},
			args: args{
				ctx: context.TODO(),
				input: usecases.SignInInputData{
					UserID: testUserProfile.UserID(),
				},
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			t.Cleanup(func() { ctrl.Finish() })

			u := interactors.NewSign(cfg, &entities.Repository{
				DynamoDB:    NewMockDynamoDB(),
				UserProfile: tt.fields.newUserProfileRepository(ctrl),
				UserSession: tt.fields.newUserSessionRepository(ctrl),
			})
			got, err := u.In(tt.args.ctx, tt.args.input)
			if !tt.wantErr(t, err) {
				return
			}
			if err == nil {
				assert.NotEmpty(t, got.UserSession)
				assert.Equal(t, tt.want.UserSession.UserID(), got.UserSession.UserID())
				assert.NotEmpty(t, got.UserSession.SessionID())
			}
		})
	}
}

func Test_signImpl_Up(t *testing.T) {
	type fields struct {
		newRegisteredUserRepository func(ctrl *gomock.Controller) entities.RegisteredUserRepository
		newUserProfileRepository    func(ctrl *gomock.Controller) entities.UserProfileRepository
	}
	type args struct {
		ctx   context.Context
		input usecases.SignUpInputData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    usecases.SignUpOutputData
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			fields: fields{
				newRegisteredUserRepository: func(ctrl *gomock.Controller) entities.RegisteredUserRepository {
					m := entities_mock.NewMockRegisteredUserRepository(ctrl)
					m.EXPECT().Create(gomock.Any(), gomock.AssignableToTypeOf(&entities.RegisteredUser{})).Return(nil)
					return m
				},
				newUserProfileRepository: func(ctrl *gomock.Controller) entities.UserProfileRepository {
					m := entities_mock.NewMockUserProfileRepository(ctrl)
					m.EXPECT().Create(gomock.Any(), gomock.AssignableToTypeOf(&entities.UserProfile{})).Return(nil)
					return m
				},
			},
			args: args{
				ctx: context.TODO(),
				input: usecases.SignUpInputData{
					Sign: "dummy-user-name",
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "failure",
			fields: fields{
				newRegisteredUserRepository: func(ctrl *gomock.Controller) entities.RegisteredUserRepository {
					m := entities_mock.NewMockRegisteredUserRepository(ctrl)
					m.EXPECT().Create(gomock.Any(), gomock.AssignableToTypeOf(&entities.RegisteredUser{})).Return(nil)
					return m
				},
				newUserProfileRepository: func(ctrl *gomock.Controller) entities.UserProfileRepository {
					m := entities_mock.NewMockUserProfileRepository(ctrl)
					m.EXPECT().Create(gomock.Any(), gomock.AssignableToTypeOf(&entities.UserProfile{})).Return(derrors.New("unknown error"))
					return m
				},
			},
			args: args{
				ctx: context.TODO(),
				input: usecases.SignUpInputData{
					Sign: "dummy-user-name",
				},
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			t.Cleanup(func() { ctrl.Finish() })

			u := interactors.NewSign(cfg, &entities.Repository{
				DynamoDB:       NewMockDynamoDB(),
				RegisteredUser: tt.fields.newRegisteredUserRepository(ctrl),
				UserProfile:    tt.fields.newUserProfileRepository(ctrl),
			})
			got, err := u.Up(tt.args.ctx, tt.args.input)
			if !tt.wantErr(t, err) {
				return
			}
			if err == nil {
				assert.NotEmpty(t, got.RegisteredUser)
				assert.NotEmpty(t, got.RegisteredUser.UserID())
				assert.NotEmpty(t, got.RegisteredUser.AccessKey())
				assert.False(t, got.RegisteredUser.CreatedAt().IsZero())
			}
		})
	}
}
