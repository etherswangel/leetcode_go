package leetcode

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	newHead := &ListNode{0, head}
	pre, cur := newHead, head

	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}

	return newHead.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
