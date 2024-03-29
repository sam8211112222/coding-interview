package Leetcode

import "testing"

func Test_buyAndSellStock1(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{prices: []int{7, 1, 5, 3, 6, 4}}, want: 5},
		{name: "case2", args: args{prices: []int{7, 6, 4, 3, 1}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buyAndSellStock1(tt.args.prices); got != tt.want {
				t.Errorf("buyAndSellStock1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buyAndSellStock2(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{prices: []int{7, 1, 5, 3, 6, 4}}, want: 7},
		{name: "case2", args: args{prices: []int{1, 2, 3, 4, 5}}, want: 4},
		{name: "case3", args: args{prices: []int{7, 6, 4, 3, 1}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buyAndSellStock2(tt.args.prices); got != tt.want {
				t.Errorf("buyAndSellStock2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buyAndSellStock3(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{prices: []int{7, 1, 5, 3, 6, 4}}, want: 7},
		{name: "case2", args: args{prices: []int{1, 2, 3, 4, 5}}, want: 4},
		{name: "case3", args: args{prices: []int{7, 6, 4, 3, 1}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buyAndSellStock2(tt.args.prices); got != tt.want {
				t.Errorf("buyAndSellStock2() = %v, want %v", got, tt.want)
			}
		})
	}
}
