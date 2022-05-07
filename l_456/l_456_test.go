package leetcode

import "testing"

func Test_find132pattern(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{3, 1, 4, 2}}, true},
		{"2", args{[]int{-1, 3, 2, 0}}, true},
		{"3", args{[]int{1, 2, 3}}, false},
		{"4", args{[]int{1, 2, 3, 4}}, false},
		{"5", args{[]int{1, 2, 3, 4, 5}}, false},
		{"6", args{[]int{1, 3, 0, 2, 5}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find132pattern(tt.args.nums); got != tt.want {
				t.Errorf("find132pattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
