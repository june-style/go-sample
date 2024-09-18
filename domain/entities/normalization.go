package entities

// MinMaxNormalization 正規化を求める
type MinMaxNormalization struct {
	min, max float64
}

func NewMinMaxNormalization(values ...float64) *MinMaxNormalization {
	min, max := MinMax(values...)
	return &MinMaxNormalization{
		min: min,
		max: max,
	}
}

func (n *MinMaxNormalization) Get(value float64) (float64, error) {
	if n.min > value {
		return 0, ErrExceedsMinimumValue
	}
	if n.max < value {
		return 1, ErrExceedsMaximumValue
	}
	return (value - n.min) / (n.max - n.min), nil
}

// ZScoreNormalization 標準化を求める
type ZScoreNormalization struct {
	min, max     float64
	mean, stddev float64
}

func NewZScoreNormalization(values ...float64) *ZScoreNormalization {
	min, max := MinMax(values...)
	mean, stddev := MeanStandardDeviation(values...)
	return &ZScoreNormalization{
		min: min, max: max,
		mean: mean, stddev: stddev,
	}
}

func (n *ZScoreNormalization) Get(value float64) (float64, error) {
	if n.min > value {
		return (n.min - n.mean) / n.stddev, ErrExceedsMinimumValue
	}
	if n.max < value {
		return (n.max - n.mean) / n.stddev, ErrExceedsMaximumValue
	}
	return (value - n.mean) / n.stddev, nil
}
