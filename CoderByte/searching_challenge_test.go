package CoderByte

import "testing"

func TestSearchingChallenge(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test case1", args: args{str: "Hello Apple Pie"}, want: "Hello"},
		{name: "test case2", args: args{str: "No words"}, want: "-1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchingChallenge(tt.args.str); got != tt.want {
				t.Errorf("SearchingChallenge() = %v, want %v", got, tt.want)
			}
		})
	}
}
