package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1",
			args{[]int{1, 1, 2}},
			[]int{1, 1, 2},
		},
		{"2",
			args{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
			[]int{0, 0, 1, 1, 2, 2, 3, 3, 4},
		},
		{"3",
			args{[]int{}},
			[]int{},
		},
		{"4",
			args{[]int{1}},
			[]int{1},
		},
		{"5",
			args{[]int{1, 1}},
			[]int{1, 1},
		},
		{"6",
			args{[]int{1, 1, 2}},
			[]int{1, 1, 2},
		},
		{"7",
			args{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}},
			[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); !reflect.DeepEqual(tt.args.nums[:got], tt.want) {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			} else {
				fmt.Println(tt.name, tt.args.nums[:got], "pass")
			}
		})
	}
}
