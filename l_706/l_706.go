package leetcode

type MyHashMap struct {
	data []*HashNode
	size int
}

type HashNode struct {
	key  int
	val  int
	next *HashNode
}

func Constructor() MyHashMap {
	return MyHashMap{data: make([]*HashNode, 100000), size: 100000}
}

func (this *MyHashMap) Put(key int, value int) {
	pos := this.Hash(key)
	node := this.data[pos]
	if node != nil {
		if node.key == key {
			node.val = value
			return
		}
		for node.next != nil {
			node = node.next
			if node.key == key {
				node.val = value
				return
			}
		}
		node.next = &HashNode{key: key, val: value, next: nil}
	} else {
		this.data[pos] = &HashNode{key, value, nil}
		return
	}
}

func (this *MyHashMap) Get(key int) int {
	pos := this.Hash(key)
	node := this.data[pos]
	for node != nil {
		if node.key == key {
			return node.val
		} else {
			node = node.next
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	pos := this.Hash(key)
	node := this.data[pos]
	var prev *HashNode = nil
	for node != nil {
		if node.key == key {
			if prev == nil {
				this.data[pos] = node.next
			} else {
				prev.next = node.next
			}
			return
		} else {
			prev = node
			node = node.next
		}
	}
}

func (this *MyHashMap) Hash(key int) int {
	return key % this.size
}
