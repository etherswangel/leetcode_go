package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	less, greater := new(ListNode), new(ListNode)
	lessHead, greaterHead := less, greater
	for head != nil {
		if head.Val < x {
			less.Next = head
			less = less.Next
		} else {
			greater.Next = head
			greater = greater.Next
		}
		head = head.Next
	}
	greater.Next = nil
	less.Next = greaterHead.Next
	return lessHead.Next
}
