package leetcode

type BSTIterator struct {
	Node     *TreeNode
	NextNode *BSTIterator
}

func Constructor(root *TreeNode) BSTIterator {
	stack := []*TreeNode{}
	res := BSTIterator{}
	iterator := &res

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		iterator.NextNode = &BSTIterator{root, nil}
		iterator = iterator.NextNode

		root = root.Right
	}
	return res
}

func (this *BSTIterator) Next() int {
	*this = *this.NextNode
	return this.Node.Val
}

func (this *BSTIterator) HasNext() bool {
	if this.NextNode == nil {
		return false
	}
	return true
}
