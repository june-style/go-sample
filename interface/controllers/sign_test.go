package controllers_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/june-style/go-sample/application/usecases"
	usecases_mock "github.com/june-style/go-sample/application/usecases/mock"
	"github.com/june-style/go-sample/domain/dcontext"
	"github.com/june-style/go-sample/domain/derrors"
	"github.com/june-style/go-sample/domain/entities"
	"github.com/june-style/go-sample/framework/protocol/pb"
	"github.com/june-style/go-sample/interface/controllers"
	"github.com/stretchr/testify/assert"
)

func Test_signImpl_In(t *testing.T) {
	var testRegisteredUser = entities.NewRegisteredUser(
		entities.GenUUID(),
		entities.GenXID(),
		time.Date(2024, 12, 25, 0, 0, 0, 1, time.Local),
	)
	var testUserSession, _ = entities.CreateUserSession(
		testRegisteredUser.UserID(),
		"dummy-secret-salt",
	)
	type fields struct {
		newSign func(ctrl *gomock.Controller) usecases.Sign
	}
	type args struct {
		ctx context.Context
		req *pb.SignInRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.SignInResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			fields: fields{
				newSign: func(ctrl *gomock.Controller) usecases.Sign {
					m := usecases_mock.NewMockSign(ctrl)
					m.EXPECT().In(gomock.Any(), usecases.SignInInputData{
						UserID: testRegisteredUser.UserID(),
					}).Return(usecases.SignInOutputData{
						UserSession: testUserSession,
					}, nil)
					return m
				},
			},
			args: args{
				ctx: dcontext.SetAuthenticatedUser(context.TODO(), testRegisteredUser),
				req: &pb.SignInRequest{},
			},
			want: &pb.SignInResponse{
				SessionId: testUserSession.SessionID(),
			},
			wantErr: assert.NoError,
		},
		{
			name: "failure",
			fields: fields{
				newSign: func(ctrl *gomock.Controller) usecases.Sign {
					m := usecases_mock.NewMockSign(ctrl)
					m.EXPECT().In(gomock.Any(), usecases.SignInInputData{
						UserID: "",
					}).Return(usecases.SignInOutputData{
						UserSession: nil,
					}, derrors.New("unknown error"))
					return m
				},
			},
			args: args{
				ctx: context.TODO(),
				req: &pb.SignInRequest{},
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			t.Cleanup(func() { ctrl.Finish() })

			c := controllers.NewSign(&usecases.UseCase{
				Sign: tt.fields.newSign(ctrl),
			})
			got, err := c.In(tt.args.ctx, tt.args.req)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_signImpl_Up(t *testing.T) {
	var testRegisteredUser = entities.NewRegisteredUser(
		entities.GenUUID(),
		entities.GenXID(),
		time.Date(2024, 12, 25, 0, 0, 0, 1, time.Local),
	)
	type fields struct {
		newSign func(ctrl *gomock.Controller) usecases.Sign
	}
	type args struct {
		ctx context.Context
		req *pb.SignUpRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.SignUpResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			fields: fields{
				newSign: func(ctrl *gomock.Controller) usecases.Sign {
					m := usecases_mock.NewMockSign(ctrl)
					m.EXPECT().Up(gomock.Any(), usecases.SignUpInputData{
						Sign: "dummy-sign",
					}).Return(usecases.SignUpOutputData{
						RegisteredUser: testRegisteredUser,
					}, nil)
					return m
				},
			},
			args: args{
				ctx: context.TODO(),
				req: &pb.SignUpRequest{
					Sign: "dummy-sign",
				},
			},
			want: &pb.SignUpResponse{
				AccessKey: testRegisteredUser.AccessKey(),
			},
			wantErr: assert.NoError,
		},
		{
			name: "failure",
			fields: fields{
				newSign: func(ctrl *gomock.Controller) usecases.Sign {
					m := usecases_mock.NewMockSign(ctrl)
					m.EXPECT().Up(gomock.Any(), usecases.SignUpInputData{
						Sign: "",
					}).Return(usecases.SignUpOutputData{
						RegisteredUser: nil,
					}, derrors.New("unknown error"))
					return m
				},
			},
			args: args{
				ctx: context.TODO(),
				req: &pb.SignUpRequest{
					Sign: "",
				},
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			t.Cleanup(func() { ctrl.Finish() })

			c := controllers.NewSign(&usecases.UseCase{
				Sign: tt.fields.newSign(ctrl),
			})
			got, err := c.Up(tt.args.ctx, tt.args.req)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
