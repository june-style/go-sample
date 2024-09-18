package entities_test

import (
	"testing"

	"github.com/june-style/go-sample/domain/entities"
	"github.com/stretchr/testify/assert"
)

func Test_Sum(t *testing.T) {
	type args struct {
		values []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "success with some values",
			args: args{
				values: []float64{
					0.0,
					1.0,
					2.0,
					3.0,
					4.0,
					5.0,
				},
			},
			want: 15.0,
		},
		{
			name: "success with empty",
			args: args{
				values: []float64{},
			},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, entities.Sum(tt.args.values...))
		})
	}
}
