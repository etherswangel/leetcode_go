package leetcode

type NestedIterator struct {
	stack   []*[]*NestedInteger
	index   []int
	hasNext []bool
	next    *NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	stack, index, hasNext, next := make([]*[]*NestedInteger, 1), make([]int, 1), make([]bool, 1), nestedList[0]
	stack[0], index[0] = &nestedList, 0
	if len(*stack[0])-1 > index[0] {
		hasNext[0] = true
	}
	return &NestedIterator{stack, index, hasNext, next}
}

func (this *NestedIterator) Next() int {
	res := *this.next

	for len(this.hasNext) != 0 && !this.hasNext[len(this.hasNext)-1] {
		this.stack = this.stack[:len(this.stack)-1]
		this.index = this.index[:len(this.index)-1]
		this.hasNext = this.hasNext[:len(this.hasNext)-1]
	}
	if len(this.hasNext) == 0 {
		this.next = nil
	} else {
		this.index[len(this.index)-1]++
		this.next = (*this.stack[len(this.stack)-1])[this.index[len(this.index)-1]]
		if len(*this.stack[len(this.stack)-1])-1 == this.index[len(this.index)-1] {
			this.hasNext[len(this.hasNext)-1] = false
		}
	}
	return res.GetInteger()
}

func (this *NestedIterator) HasNext() bool {
	for this.next != nil && !(this.next.IsInteger()) {
		list := this.next.GetList()
		if len(list) != 0 {
			this.stack = append(this.stack, &list)
			this.index = append(this.index, 0)
			if len(list) > 1 {
				this.hasNext = append(this.hasNext, true)
			} else {
				this.hasNext = append(this.hasNext, false)
			}
			this.next = list[0]
		} else {
			for len(this.hasNext) != 0 && !this.hasNext[len(this.hasNext)-1] {
				this.stack = this.stack[:len(this.stack)-1]
				this.index = this.index[:len(this.index)-1]
				this.hasNext = this.hasNext[:len(this.hasNext)-1]
			}
			if len(this.hasNext) == 0 {
				this.next = nil
			} else {
				this.index[len(this.index)-1]++
				this.next = (*this.stack[len(this.stack)-1])[this.index[len(this.index)-1]]
				if len(*this.stack[len(this.stack)-1])-1 == this.index[len(this.index)-1] {
					this.hasNext[len(this.hasNext)-1] = false
				}
			}
		}
	}
	if this.next == nil {
		return false
	}
	return true
}

///////////////////////

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool { return false }

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int { return 0 }

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger { return make([]*NestedInteger, 0) }
