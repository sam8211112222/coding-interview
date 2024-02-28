package CoderByte

import (
	"strings"
	"unicode"
)

// SearchingChallenge finds the first word with the greatest number of repeated letters
func SearchingChallenge(str string) string {
	fields := strings.Fields(str)
	var result string
	var maxRepeat int
	for _, field := range fields {
		var runeCount = make(map[rune]int)
		for _, char := range field {
			runeCount[unicode.ToLower(char)]++
			if runeCount[unicode.ToLower(char)] > 1 && runeCount[unicode.ToLower(char)] > maxRepeat {
				maxRepeat = runeCount[unicode.ToLower(char)]
				result = field
			}
		}
	}
	if result == "" {
		result = "-1"
	}
	return result
}
