package stringutil

import "strings"

func PaddingLeft(s string, padWith string, overallLen int) string {
	if overallLen < len(s) {
		return s
	}

	res := paddingString(s, padWith, overallLen) + s
	return res[(len(res) - overallLen):]
}

func PaddingRight(s string, padWith string, overallLen int) string {
	if overallLen < len(s) {
		return s
	}

	res := s + paddingString(s, padWith, overallLen)

	return res[:overallLen]
}

func paddingString(s string, padWith string, overallLen int) string {
	padCount := 1 + (overallLen-len(s)) / len(padWith)
	if padCount < 0 {
		padCount = 0
	}
	return strings.Repeat(padWith, padCount)
}
