package stringutil

import "unicode"

// UCFirst upper case first symbol in string
func UCFirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToUpper(v))
		return u + str[len(u):]
	}

	return ""
}
