package leetcode

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	return false
}

func (this *Iterator) next() int {
	return 0
}

type PeekingIterator struct {
	iter     *Iterator
	nextElem int
	haveNext bool
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		iter:     iter,
		haveNext: iter.hasNext(),
		nextElem: iter.next(),
	}
}

func (this *PeekingIterator) hasNext() bool {
	return this.haveNext
}

func (this *PeekingIterator) next() int {
	elem := this.nextElem
	this.haveNext = this.iter.hasNext()
	this.nextElem = this.iter.next()
	return elem
}

func (this *PeekingIterator) peek() int {
	return this.nextElem
}
