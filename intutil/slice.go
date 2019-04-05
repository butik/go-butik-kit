package intutil

func ReverseInts(input []int) []int {
	l := len(input)
	reversed := make([]int, l)
	for i := l-1; i >= 0; i-- {
		reversed[l-i-1] = input[i]
	}
	return reversed
}
