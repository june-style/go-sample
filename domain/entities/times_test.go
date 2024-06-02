package entities_test

import (
	"testing"
	"time"

	"github.com/june-style/go-sample/domain/entities"
	"github.com/stretchr/testify/assert"
)

type testPeriod struct {
	beginAt, endAt time.Time
}

func (t testPeriod) BeginAt() time.Time {
	return t.beginAt
}

func (t testPeriod) EndAt() time.Time {
	return t.endAt
}

func Test_IsTerm(t *testing.T) {
	type args struct {
		now    time.Time
		period entities.Period
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success with values between",
			args: args{
				now: time.Date(2024, 12, 25, 23, 59, 59, 999999999, time.Local),
				period: testPeriod{
					beginAt: time.Date(2024, 12, 25, 0, 0, 0, 0, time.Local),
					endAt:   time.Date(2024, 12, 26, 0, 0, 0, 0, time.Local),
				},
			},
			want: true,
		},
		{
			name: "success with begin value",
			args: args{
				now: time.Date(2024, 12, 25, 0, 0, 0, 0, time.Local),
				period: testPeriod{
					beginAt: time.Date(2024, 12, 25, 0, 0, 0, 0, time.Local),
					endAt:   time.Date(2024, 12, 26, 0, 0, 0, 0, time.Local),
				},
			},
			want: true,
		},
		{
			name: "failure with end value",
			args: args{
				now: time.Date(2024, 12, 26, 0, 0, 0, 0, time.Local),
				period: testPeriod{
					beginAt: time.Date(2024, 12, 25, 0, 0, 0, 0, time.Local),
					endAt:   time.Date(2024, 12, 26, 0, 0, 0, 0, time.Local),
				},
			},
			want: false,
		},
		{
			name: "failure with lower than begin value",
			args: args{
				now: time.Date(2024, 12, 24, 0, 0, 0, 0, time.Local),
				period: testPeriod{
					beginAt: time.Date(2024, 12, 25, 0, 0, 0, 0, time.Local),
					endAt:   time.Date(2024, 12, 26, 0, 0, 0, 0, time.Local),
				},
			},
			want: false,
		},
		{
			name: "failure with higher than end value",
			args: args{
				now: time.Date(2024, 12, 26, 0, 0, 0, 0, time.Local),
				period: testPeriod{
					beginAt: time.Date(2024, 12, 25, 0, 0, 0, 0, time.Local),
					endAt:   time.Date(2024, 12, 26, 0, 0, 0, 0, time.Local),
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, entities.IsTerm(tt.args.now, tt.args.period))
		})
	}
}
