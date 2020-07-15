package word

import (
	"strings"
	"unicode"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func UndersocreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

func UndersoreToLowerCamelCase(s string) string {
	s = UndersocreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}
