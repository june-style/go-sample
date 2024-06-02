package entities_test

import (
	"testing"

	"github.com/june-style/go-sample/domain/entities"
	"github.com/stretchr/testify/assert"
)

func Test_MeanStandardDeviation(t *testing.T) {
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
					100, 110, 70, 50, 120,
				},
			},
			want:  90,
			want1: 26.076809620810597,
		},
		{
			name: "success with example 2",
			args: args{
				values: []float64{
					1.9, 2.5, 2.5, 3.5, 3.5, 4.0,
				},
			},
			want:  2.9833333333333334,
			want1: 0.7312470322826766,
		},
		{
			name: "success with example 3",
			args: args{
				values: []float64{
					0, 0.0,
				},
			},
			want:  0.0,
			want1: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := entities.MeanStandardDeviation(tt.args.values...)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
