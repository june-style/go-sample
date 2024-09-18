package dcontext_test

import (
	"context"
	"fmt"
	"maps"
	"testing"

	"github.com/june-style/go-sample/domain/dcontext"
	"github.com/stretchr/testify/assert"
)

var (
	testHeaders = []dcontext.Header{
		dcontext.HeaderApplicationKey,
		dcontext.HeaderAccessKey,
		dcontext.HeaderSessionID,
		dcontext.HeaderRequestID,
		dcontext.HeaderUserAgent,
	}
	testMapHeaders = func() map[string]string {
		var headers = make(map[string]string, len(testHeaders))
		for _, header := range testHeaders {
			headers[header.String()] = fmt.Sprintf("dummy-%s", header)
		}
		return headers
	}()
	testMapXHeaders = func() map[string]string {
		var headers = make(map[string]string, len(testMapHeaders)-1)
		maps.Copy(headers, testMapHeaders)
		delete(headers, dcontext.HeaderUserAgent.String())
		return headers
	}()
)

func Test_GetHeader(t *testing.T) {
	type args struct {
		ctx    context.Context
		header dcontext.Header
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success with x-application-key",
			args: args{
				ctx:    dcontext.SetMetadata(context.TODO(), testMetadata),
				header: dcontext.HeaderApplicationKey,
			},
			want:    fmt.Sprintf("dummy-%s", dcontext.HeaderApplicationKey),
			wantErr: assert.NoError,
		},
		{
			name: "success with x-access-key",
			args: args{
				ctx:    dcontext.SetMetadata(context.TODO(), testMetadata),
				header: dcontext.HeaderAccessKey,
			},
			want:    fmt.Sprintf("dummy-%s", dcontext.HeaderAccessKey),
			wantErr: assert.NoError,
		},
		{
			name: "success with x-session-id",
			args: args{
				ctx:    dcontext.SetMetadata(context.TODO(), testMetadata),
				header: dcontext.HeaderSessionID,
			},
			want:    fmt.Sprintf("dummy-%s", dcontext.HeaderSessionID),
			wantErr: assert.NoError,
		},
		{
			name: "success with x-request-id",
			args: args{
				ctx:    dcontext.SetMetadata(context.TODO(), testMetadata),
				header: dcontext.HeaderRequestID,
			},
			want:    fmt.Sprintf("dummy-%s", dcontext.HeaderRequestID),
			wantErr: assert.NoError,
		},
		{
			name: "success with user-agent",
			args: args{
				ctx:    dcontext.SetMetadata(context.TODO(), testMetadata),
				header: dcontext.HeaderUserAgent,
			},
			want:    fmt.Sprintf("dummy-%s", dcontext.HeaderUserAgent),
			wantErr: assert.NoError,
		},
		{
			name: "failure with empty header",
			args: args{
				ctx:    context.TODO(),
				header: dcontext.HeaderUserAgent,
			},
			want:    "",
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dcontext.GetHeader(tt.args.ctx, tt.args.header)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_GetHeaders(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success with test data",
			args: args{
				ctx: dcontext.SetMetadata(context.TODO(), testMetadata),
			},
			want:    testMapXHeaders,
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
			got, err := dcontext.GetHeaders(tt.args.ctx)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_GetUserAgent(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success with test data",
			args: args{
				ctx: dcontext.SetMetadata(context.TODO(), testMetadata),
			},
			want:    testMapHeaders[dcontext.HeaderUserAgent.String()],
			wantErr: assert.NoError,
		},
		{
			name: "failure with empty data",
			args: args{
				ctx: context.TODO(),
			},
			want:    "",
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dcontext.GetUserAgent(tt.args.ctx)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
