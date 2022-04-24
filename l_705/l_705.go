package leetcode

type MyHashSet struct {
	data map[int]bool
}

func Constructor() MyHashSet {
	return MyHashSet{data: make(map[int]bool)}
}

func (this *MyHashSet) Add(key int) {
	this.data[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.data[key] = false
}

func (this *MyHashSet) Contains(key int) bool {
	return this.data[key]
}
