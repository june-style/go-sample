package entities_test

import (
	"testing"

	"github.com/june-style/go-sample/domain/entities"
	"github.com/stretchr/testify/assert"
)

func Test_MinMaxNormalization_Get(t *testing.T) {
	type fields struct {
		values []float64
	}
	type args struct {
		value float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success with minimum value",
			fields: fields{
				values: []float64{
					65, 80, 80, 100, 95,
				},
			},
			args: args{
				value: 65,
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "success with maximum value",
			fields: fields{
				values: []float64{
					65, 80, 80, 100, 95,
				},
			},
			args: args{
				value: 100,
			},
			want:    1,
			wantErr: assert.NoError,
		},
		{
			name: "failure with exceeds minimum value",
			fields: fields{
				values: []float64{
					65, 80, 80, 100, 95,
				},
			},
			args: args{
				value: 50,
			},
			want: 0,
			wantErr: func(t assert.TestingT, err error, msgAndArgs ...any) bool {
				return assert.ErrorIs(t, err, entities.ErrExceedsMinimumValue, msgAndArgs...)
			},
		},
		{
			name: "failure with exceeds maximum value",
			fields: fields{
				values: []float64{
					65, 80, 80, 100, 95,
				},
			},
			args: args{
				value: 120,
			},
			want: 1,
			wantErr: func(t assert.TestingT, err error, msgAndArgs ...any) bool {
				return assert.ErrorIs(t, err, entities.ErrExceedsMaximumValue, msgAndArgs...)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := entities.NewMinMaxNormalization(tt.fields.values...)
			got, err := n.Get(tt.args.value)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_ZScoreNormalization_Get(t *testing.T) {
	type fields struct {
		values []float64
	}
	type args struct {
		value float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success with minimum value",
			fields: fields{
				values: []float64{
					65, 80, 80, 100, 95,
				},
			},
			args: args{
				value: 65,
			},
			want:    -1.5310636316482225,
			wantErr: assert.NoError,
		},
		{
			name: "success with maximum value",
			fields: fields{
				values: []float64{
					65, 80, 80, 100, 95,
				},
			},
			args: args{
				value: 100,
			},
			want:    1.2893167424406085,
			wantErr: assert.NoError,
		},
		{
			name: "failure with exceeds minimum value",
			fields: fields{
				values: []float64{
					65, 80, 80, 100, 95,
				},
			},
			args: args{
				value: 50,
			},
			want: -1.5310636316482225,
			wantErr: func(t assert.TestingT, err error, msgAndArgs ...any) bool {
				return assert.ErrorIs(t, err, entities.ErrExceedsMinimumValue, msgAndArgs...)
			},
		},
		{
			name: "failure with exceeds maximum value",
			fields: fields{
				values: []float64{
					65, 80, 80, 100, 95,
				},
			},
			args: args{
				value: 120,
			},
			want: 1.2893167424406085,
			wantErr: func(t assert.TestingT, err error, msgAndArgs ...any) bool {
				return assert.ErrorIs(t, err, entities.ErrExceedsMaximumValue, msgAndArgs...)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := entities.NewZScoreNormalization(tt.fields.values...)
			got, err := n.Get(tt.args.value)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
