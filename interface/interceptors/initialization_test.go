package interceptors_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/june-style/go-sample/domain/dcontext"
	"github.com/june-style/go-sample/domain/services"
	services_mock "github.com/june-style/go-sample/domain/services/mock"
	"github.com/june-style/go-sample/interface/interceptors"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func Test_Initialization(t *testing.T) {
	var testDate = time.Date(2024, 12, 25, 0, 0, 0, 1, time.Local)
	var testUnaryInfo = &grpc.UnaryServerInfo{
		FullMethod: "TestService.UnaryMethod",
	}
	var testUnaryHandler = func(ctx context.Context, req any) (any, error) {
		return dcontext.GetTime(ctx)
	}
	type fields struct {
		newTimer func(ctrl *gomock.Controller) services.Timer
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
		want    time.Time
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			fields: fields{
				newTimer: func(ctrl *gomock.Controller) services.Timer {
					m := services_mock.NewMockTimer(ctrl)
					m.EXPECT().SetNow(gomock.Any()).Return(dcontext.SetTime(context.TODO(), testDate))
					return m
				},
			},
			args: args{
				ctx:     context.TODO(),
				req:     nil,
				info:    testUnaryInfo,
				handler: testUnaryHandler,
			},
			want:    testDate,
			wantErr: assert.NoError,
		},
		{
			name: "failure",
			fields: fields{
				newTimer: func(ctrl *gomock.Controller) services.Timer {
					m := services_mock.NewMockTimer(ctrl)
					m.EXPECT().SetNow(gomock.Any()).Return(context.TODO())
					return m
				},
			},
			args: args{
				ctx:     context.TODO(),
				req:     nil,
				info:    testUnaryInfo,
				handler: testUnaryHandler,
			},
			want:    time.Time{},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			t.Cleanup(func() { ctrl.Finish() })

			i := interceptors.Initialization(tt.fields.newTimer(ctrl))
			got, err := i(tt.args.ctx, tt.args.req, tt.args.info, tt.args.handler)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
