package string_category

import "strings"

func strStr(haystack string, needle string) int {

	//直接使用golang函數庫
	//return strings.Index(haystack, needle)

	// 這個不夠好, 空間複雜度是O(N)
	//var m = make(map[string]int)
	//length := len(haystack) - len(needle)
	//for i := 0; i <= length; i++ {
	//	s := haystack[i : i+len(needle)]
	//	_, ok := m[s]
	//	if !ok {
	//		m[s] = i
	//	}
	//}
	//v, ok := m[needle]
	//if ok {
	//	return v
	//}
	//return -1

	// 這個在leetcodeRuntime
	//0
	//ms
	//Beats
	//100.00%
	//of users with Go
	// 但 memory還是不好
	length := len(haystack) - len(needle)
	for i := 0; i <= length; i++ {
		str := haystack[i : i+len(needle)]
		if strings.EqualFold(str, needle) {
			return i
		}
	}
	return -1
}
