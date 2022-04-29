package leetcode

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	root := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return root
}

func reverseList1(head *ListNode) *ListNode {
	var behind *ListNode
	for head != nil {
		next := head.Next
		head.Next = behind
		behind = head
		head = next
	}
	return behind
}

type ListNode struct {
	Val  int
	Next *ListNode
}
