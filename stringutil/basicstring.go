package stringutil

import (
	"unicode"
	"unicode/utf8"
)

func IsEmpty(val string) bool {
	return val == ""
}

func IsUpperFirst(val string) bool {
	r, _ := utf8.DecodeRuneInString(val)
	return unicode.IsUpper(r)
}

func IsUpperLast(val string) bool {
	r, _ := utf8.DecodeLastRuneInString(val)
	return unicode.IsUpper(r)
}

func IsLowerFirst(val string) bool {
	r, _ := utf8.DecodeRuneInString(val)
	return unicode.IsLower(r)
}

func IsLowerLast(val string) bool {
	r, _ := utf8.DecodeLastRuneInString(val)
	return unicode.IsLower(r)
}

func StrLen(val string) int {
	return utf8.RuneCountInString(val)
}
