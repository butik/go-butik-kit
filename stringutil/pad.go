package stringutil

import "strings"

// PaddingLeft add padWith to the left of s until it reaches a length overallLen
func PaddingLeft(s string, padWith string, overallLen uint8) string {
	if int(overallLen) < len(s) {
		return s
	}

	res := paddingString(s, padWith, overallLen) + s
	return res[(len(res) - int(overallLen)):]
}

// PaddingRight add padWith to the right of s until it reaches a length overallLen
func PaddingRight(s string, padWith string, overallLen uint8) string {
	if int(overallLen) < len(s) {
		return s
	}

	res := s + paddingString(s, padWith, overallLen)

	return res[:overallLen]
}

func paddingString(s string, padWith string, overallLen uint8) string {
	padCount := 1 + (int(overallLen)-len(s))/len(padWith)
	if padCount < 0 {
		padCount = 0
	}
	return strings.Repeat(padWith, padCount)
}
