package questions

import "testing"

func Test_compareString(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test case 1", args: args{str1: "Hello", str2: "Hello"}, want: true},
		{name: "test case 2", args: args{str1: "Hello", str2: "hello"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareString(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("compareString() = %v, want %v", got, tt.want)
			}
		})
	}
}
