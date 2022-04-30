package leetcode

type MyQueue struct {
	_q []int
}

func Constructor() MyQueue {
	return MyQueue{make([]int, 0, 100)}
}

func (this *MyQueue) Push(x int) {
	this._q = append(this._q, x)
}

func (this *MyQueue) Pop() int {
	res := this._q[0]
	this._q = this._q[1:]
	return res
}

func (this *MyQueue) Peek() int {
	return this._q[0]
}

func (this *MyQueue) Empty() bool {
	return len(this._q) == 0
}
