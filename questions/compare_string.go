package questions

import "strings"

func compareString(str1 string, str2 string) bool {

	if len(str1) != len(str2) {
		return false
	}

	m := map[string]int{}
	for _, v := range str1 {
		m[strings.ToLower(string(v))]++
	}
	for _, v := range str2 {
		m[strings.ToLower(string(v))]--
		if m[strings.ToLower(string(v))] < 0 {
			return false
		}
	}
	return true
}
