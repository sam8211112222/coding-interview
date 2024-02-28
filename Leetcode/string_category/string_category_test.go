package string_category

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test case 1", args: args{
			haystack: "sadbutsad",
			needle:   "sad",
		}, want: 0},
		{name: "test case 2", args: args{
			haystack: "leetcode",
			needle:   "leeto",
		}, want: -1},
		{name: "test case 3", args: args{
			haystack: "aaaaa",
			needle:   "bba",
		}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
