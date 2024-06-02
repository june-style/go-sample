package entities

import "math"

// MeanStandardDeviation 平均値と標準偏差を求める
func MeanStandardDeviation(values ...float64) (float64, float64) {
	n := float64(len(values))
	m, s := 0.0, 0.0
	for i, x := range values {
		x -= m
		m += x / float64(i+1)
		s += float64(i) * x * x / float64(i+1)
	}
	s = math.Sqrt(s / n)
	return m, s
}
