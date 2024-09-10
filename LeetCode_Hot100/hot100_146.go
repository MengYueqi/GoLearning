package main

type Node struct {
	val        int
	next, prev *Node
}

type LRUCache struct {
	capNum     int
	keyNodeMap map[int]*Node
	root, tail *Node
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{capNum: capacity, keyNodeMap: make(map[int]*Node), root: nil}
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.keyNodeMap[key]; exists {
		// 将相关数据放到第一个
		if node == this.tail {
			this.tail = node.prev
			if this.tail == nil {
				this.tail = node
			}
			nextTemp, prevTemp := node.next, node.prev
			node.prev.next, node.next.prev = nextTemp, prevTemp
			nextTemp = this.root.next
			this.root.next = node
			node.prev = this.root
			node.next = nextTemp
		} else {
			nextTemp, prevTemp := node.next, node.prev
			node.prev.next, node.next.prev = nextTemp, prevTemp
			nextTemp = this.root.next
			this.root.next = node
			node.prev = this.root
			node.next = nextTemp
		}
		return this.keyNodeMap[key].val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	this.capNum--
	var node *Node
	if this.root == nil {
		node = &Node{val: value, prev: nil, next: nil}
		this.root = node
		this.tail = node
	} else {
		tempNext := this.root
		node = &Node{val: value, prev: nil, next: tempNext}
		this.root = node
	}
	this.keyNodeMap[key] = node
	if this.capNum < 0 {
		this.tail = this.tail.prev
		this.tail.next = nil
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
