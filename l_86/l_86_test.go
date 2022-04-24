package leetcode

import (
	"fmt"
	"testing"
)

func generateList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{
		Val: nums[0],
	}
	tmp := head
	for i := 1; i < len(nums); i++ {
		tmp.Next = &ListNode{
			Val: nums[i],
		}
		tmp = tmp.Next
	}
	return head
}

func equalList(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val, "->")
		l = l.Next
	}
	fmt.Println()
}

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1",
			args{generateList([]int{1, 4, 3, 2, 5, 2}), 3},
			generateList([]int{1, 2, 2, 4, 3, 5}),
		},
		{"2",
			args{nil, 1},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !equalList(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			} else {
				printList(got)
				printList(tt.want)
			}
		})
	}
}
