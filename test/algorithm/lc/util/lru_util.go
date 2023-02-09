package util

// LRUCache
// Least Recently Used: 最近最少使用
// LRU 是由一个 map 和一个双向链表组成的数据结构
// map 中 key 对应的 value 是双向链表的结点
// 通过dummyHead 和 dummyTail 来避免边界判断
type LRUCache struct {
	keys       map[int]*Node
	head, tail *Node
	capacity   int
}

type Node struct {
	key, value int
	prev, next *Node
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		keys:     make(map[int]*Node),
		head:     &Node{},
		tail:     &Node{},
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.keys[key]; ok {
		// move to head -> remove + add
		l.Remove(node)
		l.Add(node)
		return node.value
	}
	return -1
}

func (l *LRUCache) Put(k, v int) {
	if _, ok := l.keys[k]; ok {
		node := l.keys[k]
		node.value = v
		l.Remove(node)
		l.Add(node)
	} else {
		n := &Node{key: k, value: v}
		l.keys[k] = n
		l.Add(n)
		if len(l.keys) > l.capacity {
			tail := l.tail.prev
			l.Remove(tail)
			delete(l.keys, tail.key)
		}
	}
}

func (l *LRUCache) Add(n *Node) {
	n.next = l.head.next
	n.prev = l.head
	l.head.next.prev = n
	l.head.next = n
}

func (l *LRUCache) Remove(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}
