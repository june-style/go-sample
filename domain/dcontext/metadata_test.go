package dcontext_test

import (
	"context"
	"testing"

	"github.com/june-style/go-sample/domain/dcontext"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

var testMetadata = metadata.New(testMapHeaders)

func Test_GetMetadata(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    metadata.MD
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success with test data",
			args: args{
				ctx: dcontext.SetMetadata(context.TODO(), testMetadata),
			},
			want:    testMetadata,
			wantErr: assert.NoError,
		},
		{
			name: "failure with empty data",
			args: args{
				ctx: context.TODO(),
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dcontext.GetMetadata(tt.args.ctx)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_SetMetadata(t *testing.T) {
	type args struct {
		ctx context.Context
		md  metadata.MD
	}
	tests := []struct {
		name string
		args args
		want metadata.MD
	}{
		{
			name: "success with test data",
			args: args{
				ctx: context.TODO(),
				md:  testMetadata,
			},
			want: testMetadata,
		},
		{
			name: "success with empty data",
			args: args{
				ctx: context.TODO(),
				md:  nil,
			},
			want: metadata.MD{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := dcontext.SetMetadata(tt.args.ctx, tt.args.md)
			md, _ := dcontext.GetMetadata(ctx)
			assert.Equal(t, tt.want, md)
		})
	}
}
