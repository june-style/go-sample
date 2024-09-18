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
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_homeImpl_Get(t *testing.T) {
	var testRegisteredUser = entities.NewRegisteredUser(
		entities.GenUUID(),
		entities.GenXID(),
		time.Date(2024, 12, 25, 0, 0, 0, 1, time.Local),
	)
	var testUserProfile = entities.NewUserProfile(
		testRegisteredUser.UserID(),
		"dummy-user-name",
		testRegisteredUser.CreatedAt(),
	)
	type fields struct {
		newHome func(ctrl *gomock.Controller) usecases.Home
	}
	type args struct {
		ctx context.Context
		req *pb.HomeGetRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.HomeGetResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			fields: fields{
				newHome: func(ctrl *gomock.Controller) usecases.Home {
					m := usecases_mock.NewMockHome(ctrl)
					m.EXPECT().Get(gomock.Any(), usecases.HomeGetInputData{
						UserID: testRegisteredUser.UserID(),
					}).Return(usecases.HomeGetOutputData{
						UserProfile: testUserProfile,
					}, nil)
					return m
				},
			},
			args: args{
				ctx: dcontext.SetAuthenticatedUser(context.TODO(), testRegisteredUser),
				req: &pb.HomeGetRequest{},
			},
			want: &pb.HomeGetResponse{
				User: &pb.User{
					Id:        testUserProfile.UserID(),
					Name:      testUserProfile.Name(),
					CreatedAt: timestamppb.New(testUserProfile.CreatedAt()),
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "failure",
			fields: fields{
				newHome: func(ctrl *gomock.Controller) usecases.Home {
					m := usecases_mock.NewMockHome(ctrl)
					m.EXPECT().Get(gomock.Any(), usecases.HomeGetInputData{
						UserID: "",
					}).Return(usecases.HomeGetOutputData{
						UserProfile: nil,
					}, derrors.New("unknown error"))
					return m
				},
			},
			args: args{
				ctx: context.TODO(),
				req: &pb.HomeGetRequest{},
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			t.Cleanup(func() { ctrl.Finish() })

			c := controllers.NewHome(&usecases.UseCase{
				Home: tt.fields.newHome(ctrl),
			})
			got, err := c.Get(tt.args.ctx, tt.args.req)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
