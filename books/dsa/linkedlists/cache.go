package linkedlists

import (
	"container/list"
)

// LRUCache using container/list (Go standard library equivalent to OrderedDict)
type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type entry struct {
	key   int
	value int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		return elem.Value.(*entry).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		elem.Value.(*entry).value = value
		return
	}

	if this.list.Len() >= this.capacity {
		// Evict least recently used (at the back)
		back := this.list.Back()
		if back != nil {
			this.list.Remove(back)
			delete(this.cache, back.Value.(*entry).key)
		}
	}

	elem := this.list.PushFront(&entry{key: key, value: value})
	this.cache[key] = elem
}

// ManualLRUCache using custom DLL + Hashmap
type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

type ManualLRUCache struct {
	capacity int
	cache    map[int]*Node
	left     *Node // LRU sentinel
	right    *Node // MRU sentinel
}

func NewManualLRUCache(capacity int) *ManualLRUCache {
	left := &Node{}
	right := &Node{}
	left.next = right
	right.prev = left

	return &ManualLRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		left:     left,
		right:    right,
	}
}

func (this *ManualLRUCache) remove(node *Node) {
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
}

func (this *ManualLRUCache) insertMRU(node *Node) {
	prev, next := this.right.prev, this.right
	prev.next = node
	next.prev = node
	node.prev = prev
	node.next = next
}

func (this *ManualLRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.remove(node)
		this.insertMRU(node)
		return node.val
	}
	return -1
}

func (this *ManualLRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		this.remove(node)
		node.val = value
		this.insertMRU(node)
		return
	}

	if len(this.cache) >= this.capacity {
		// Evict LRU node (left.next)
		lru := this.left.next
		this.remove(lru)
		delete(this.cache, lru.key)
	}

	node := &Node{key: key, val: value}
	this.cache[key] = node
	this.insertMRU(node)
}
