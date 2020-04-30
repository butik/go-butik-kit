package sliceutils

func Min(n int, less func(int, int) bool) int {
	return FoldLeft(n, 0, func(acc interface{}, i int) interface{} {
		if less(acc.(int), i) {
			return i
		} else {
			return acc.(int)
		}
	}).(int)
}

func Max(n int, greater func(int, int) bool) int {
	return FoldLeft(n, 0, func(acc interface{}, i int) interface{} {
		if greater(acc.(int), i) {
			return i
		} else {
			return acc.(int)
		}
	}).(int)
}
