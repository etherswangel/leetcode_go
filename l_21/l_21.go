package leetcode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val > list2.Val {
		return mergeTwoLists(list2, list1)
	}

	p1, p2, t := list1.Next, list2, list1
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			t = t.Next
			p1 = p1.Next
		} else {
			t.Next = p2
			p2 = p2.Next
			t = t.Next
			t.Next = p1
		}
	}

	if p1 == nil {
		t.Next = p2
	}

	return list1
}

type ListNode struct {
	Val  int
	Next *ListNode
}
