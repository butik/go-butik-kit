package sliceutils

func FoldLeft(n int, init interface{}, f func(acc interface{}, i int) interface{}) interface{} {
	acc := init
	for i := 0; i < n; i++ {
		acc = f(acc, i)
	}
	return acc
}

func FoldRight(n int, init interface{}, f func(i int, acc interface{}) interface{}) interface{} {
	acc := init
	for i := n; i >= 0; i-- {
		acc = f(i, acc)
	}
	return acc
}
