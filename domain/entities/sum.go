package entities

// Sum 合計値を求める
func Sum[T Number](values ...T) T {
	a := T(0)
	for _, v := range values {
		a += v
	}
	return a
}
