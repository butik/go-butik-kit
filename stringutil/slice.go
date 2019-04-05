package stringutil

func ReverseStrings(input []string) []string {
	l := len(input)
	reversed := make([]string, l)
	for i := l-1; i >= 0; i-- {
		reversed[l-i-1] = input[i]
	}
	return reversed
}
