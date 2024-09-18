package entities

import "time"

const defaultFormat = time.RFC3339Nano

func TimeFormat(time time.Time) string {
	return time.Format(defaultFormat)
}

func Now() time.Time {
	return time.Now()
}

type Period interface {
	BeginAt() time.Time
	EndAt() time.Time
}

func IsTerm(now time.Time, period Period) bool {
	begin := period.BeginAt()
	end := period.EndAt()
	return !now.Before(begin) && !now.After(end) && !now.Equal(end)
}
