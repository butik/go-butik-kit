package stringutil

import "strings"

// PaddingLeft add padWith to the left of s until it reaches a length overallLen
func PaddingLeft(s string, padWith string, overallLen int) string {
	if overallLen < len(s) {
		return s
	}

	res := paddingString(s, padWith, overallLen) + s
	return res[(len(res) - overallLen):]
}

// PaddingRight add padWith to the right of s until it reaches a length overallLen
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
