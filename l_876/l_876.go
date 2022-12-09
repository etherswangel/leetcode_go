package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	r := head
	for r != nil && r.Next != nil {
		r = r.Next.Next
		head = head.Next
	}
	return head
}
