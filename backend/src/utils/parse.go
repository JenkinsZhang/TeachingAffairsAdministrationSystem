package utils

import "unicode"

// ok!
func OnlyLetterAndNum(s string) bool {
	for _, c := range s {
		if unicode.IsLetter(c) {
			continue
		}
		if unicode.IsNumber(c) {
			continue
		}
		return false
	}
	return true
}
