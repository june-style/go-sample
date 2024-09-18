package dcontext_test

import (
	"context"
	"testing"
	"time"

	"github.com/june-style/go-sample/domain/dcontext"
	"github.com/june-style/go-sample/domain/entities"
	"github.com/stretchr/testify/assert"
)

var now = entities.Now()

func Test_GetTime(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success with test data",
			args: args{
				ctx: dcontext.SetTime(context.TODO(), now),
			},
			want:    now,
			wantErr: assert.NoError,
		},
		{
			name: "failure with empty data",
			args: args{
				ctx: context.TODO(),
			},
			want:    time.Time{},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dcontext.GetTime(tt.args.ctx)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
