package entities_test

import (
	"testing"

	"github.com/june-style/go-sample/domain/entities"
	"github.com/stretchr/testify/assert"
)

func Test_MinMax(t *testing.T) {
	type args struct {
		values []float64
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
	}{
		{
			name: "success with example 1",
			args: args{
				values: []float64{
					1, 10, 5, 7, 0,
				},
			},
			want:  0,
			want1: 10,
		},
		{
			name: "success with example 2",
			args: args{
				values: []float64{
					3,
				},
			},
			want:  3,
			want1: 3,
		},
		{
			name: "success with example 3",
			args: args{
				values: []float64{},
			},
			want:  0,
			want1: 0,
		},
		{
			name: "success with example 3",
			args: args{
				values: nil,
			},
			want:  0,
			want1: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := entities.MinMax(tt.args.values...)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
