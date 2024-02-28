package Leetcode

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		input  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test case1", args: args{
			input:  []int{2, 7, 11, 15},
			target: 9,
		}, want: []int{0, 1}},
		{name: "test case2", args: args{
			input:  []int{3, 2, 4},
			target: 6,
		}, want: []int{1, 2}},
		{name: "test case3", args: args{
			input:  []int{3, 3},
			target: 6,
		}, want: []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.input, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threeSum(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "test case1", args: args{
			input: []int{-1, 0, 1, 2, -1, -4},
		}, want: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{name: "test case2", args: args{
			input: []int{0, 1, 1},
		}, want: [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
