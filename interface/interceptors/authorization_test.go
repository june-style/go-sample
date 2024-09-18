package interceptors_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/june-style/go-sample/domain/dcontext"
	"github.com/june-style/go-sample/domain/entities"
	"github.com/june-style/go-sample/domain/services"
	services_mock "github.com/june-style/go-sample/domain/services/mock"
	"github.com/june-style/go-sample/framework/protocol/pb"
	"github.com/june-style/go-sample/interface/interceptors"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func Test_Authorization(t *testing.T) {
	var testRegisteredUser = entities.NewRegisteredUser(
		entities.GenUUID(),
		entities.GenXID(),
		time.Date(2024, 12, 25, 0, 0, 0, 1, time.Local),
	)
	var testContext = dcontext.SetAuthenticatedUser(context.TODO(), testRegisteredUser)
	var testUnaryHandler = func(ctx context.Context, req any) (any, error) {
		return nil, nil
	}
	type fields struct {
		newAuthorizer func(ctrl *gomock.Controller) services.Authorizer
	}
	type args struct {
		ctx     context.Context
		req     any
		info    *grpc.UnaryServerInfo
		handler grpc.UnaryHandler
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			fields: fields{
				newAuthorizer: func(ctrl *gomock.Controller) services.Authorizer {
					m := services_mock.NewMockAuthorizer(ctrl)
					m.EXPECT().VerifyApplicationKey(gomock.Any()).Return(nil)
					m.EXPECT().VerifyAccessKey(gomock.Any()).Return(testContext, nil)
					m.EXPECT().VerifySession(gomock.Any()).Return(nil)
					return m
				},
			},
			args: args{
				ctx: context.TODO(),
				req: &pb.HomeGetRequest{},
				info: &grpc.UnaryServerInfo{
					FullMethod: pb.HomeService_Get_FullMethodName,
				},
				handler: testUnaryHandler,
			},
			want:    nil,
			wantErr: assert.NoError,
		},
		{
			name: "failure",
			fields: fields{
				newAuthorizer: func(ctrl *gomock.Controller) services.Authorizer {
					m := services_mock.NewMockAuthorizer(ctrl)
					return m
				},
			},
			args: args{
				ctx:     context.TODO(),
				req:     nil,
				info:    nil,
				handler: testUnaryHandler,
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			t.Cleanup(func() { ctrl.Finish() })

			i := interceptors.Authorization(tt.fields.newAuthorizer(ctrl))
			got, err := i(tt.args.ctx, tt.args.req, tt.args.info, tt.args.handler)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_GetMethodOptionAdmin(t *testing.T) {
	type args struct {
		req  any
		info *grpc.UnaryServerInfo
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.Admin
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success with /api.HomeService/Get",
			args: args{
				req: &pb.HomeGetRequest{},
				info: &grpc.UnaryServerInfo{
					FullMethod: pb.HomeService_Get_FullMethodName,
				},
			},
			want: &pb.Admin{
				DisableToAuthAccessKey: false,
				DisableToAuthSessionId: false,
			},
			wantErr: assert.NoError,
		},
		{
			name: "success with /api.SignService/In",
			args: args{
				req: &pb.SignInRequest{},
				info: &grpc.UnaryServerInfo{
					FullMethod: pb.SignService_In_FullMethodName,
				},
			},
			want: &pb.Admin{
				DisableToAuthAccessKey: false,
				DisableToAuthSessionId: true,
			},
			wantErr: assert.NoError,
		},
		{
			name: "success with /api.SignService/Up",
			args: args{
				req: &pb.SignUpRequest{},
				info: &grpc.UnaryServerInfo{
					FullMethod: pb.SignService_Up_FullMethodName,
				},
			},
			want: &pb.Admin{
				DisableToAuthAccessKey: true,
				DisableToAuthSessionId: true,
			},
			wantErr: assert.NoError,
		},
		{
			name: "failure",
			args: args{
				req:  nil,
				info: nil,
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := interceptors.GetMethodOptionAdmin(tt.args.req, tt.args.info)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want.String(), got.String())
		})
	}
}
