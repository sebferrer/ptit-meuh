package util

import "strings"

func ReplaceSpecialChars(s string) string {
	s = strings.ReplaceAll(s, "À", "A")
	s = strings.ReplaceAll(s, "Â", "A")
	s = strings.ReplaceAll(s, "Ä", "A")
	s = strings.ReplaceAll(s, "É", "E")
	s = strings.ReplaceAll(s, "È", "E")
	s = strings.ReplaceAll(s, "Ê", "E")
	s = strings.ReplaceAll(s, "Ë", "E")
	s = strings.ReplaceAll(s, "Ï", "I")
	s = strings.ReplaceAll(s, "Î", "I")
	s = strings.ReplaceAll(s, "Ô", "O")
	s = strings.ReplaceAll(s, "Ö", "O")
	s = strings.ReplaceAll(s, "Ù", "U")
	s = strings.ReplaceAll(s, "Û", "U")
	s = strings.ReplaceAll(s, "Ü", "U")
	s = strings.ReplaceAll(s, "Ÿ", "Y")
	s = strings.ReplaceAll(s, "Ç", "C")

	return s
}
