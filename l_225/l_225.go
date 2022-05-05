package leetcode

type MyStack struct {
	stack []int
}

func Constructor() MyStack {
	return MyStack{make([]int, 0)}
}

func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyStack) Pop() int {
	res := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return res
}

func (this *MyStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.stack) == 0
}
