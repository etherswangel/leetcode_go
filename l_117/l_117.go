package leetcode

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	size, stack := 1, make([]*Node, 0)
	stack = append(stack, root)

	for size != 0 {
		for i := 0; i < size; i++ {
			if i == size-1 {
				stack[i].Next = nil
			} else {
				stack[i].Next = stack[i+1]
			}
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		size, stack = len(stack)-size, stack[size:]
	}

	return root
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
