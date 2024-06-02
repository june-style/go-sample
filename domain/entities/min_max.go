package entities

// MinMax 最小値と最大値を求める
func MinMax[T Number](values ...T) (T, T) {
	if len(values) == 0 {
		return T(0), T(0)
	}
	if len(values) == 1 {
		return T(values[0]), T(values[0])
	}
	min, max := values[0], values[0]
	for _, v := range values[1:] {
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	return min, max
}
