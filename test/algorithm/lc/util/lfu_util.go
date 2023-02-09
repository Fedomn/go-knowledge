package util

// LFUCache
// Least Frequently Used: 最不经常使用
// 它只关心最小频次，不需要排序。因此保存最小频次，淘汰时读取这个最小值能找到要删除的结点
type LFUCache struct {
	minFreq  int
	keys     map[int]*Node2
	freqs    map[int]*List2
	capacity int
}

func ConstructorLFU(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		keys:     make(map[int]*Node2, 0),
		freqs:    make(map[int]*List2, 0),
		minFreq:  0,
	}
}

func (l *LFUCache) Get(key int) int {
	if n, ok := l.keys[key]; ok {
		freq := n.freq
		l.freqs[freq].Remove(n)
		// clean freq if possible
		if l.freqs[freq].size == 0 {
			delete(l.freqs, freq)
			if freq == l.minFreq {
				l.minFreq++
			}
		}
		// 插入节点至freq+1
		n.freq++
		if _, ok := l.freqs[freq+1]; ok {
			l.freqs[freq+1].Add(n)
		} else {
			nl := NewList2()
			nl.Add(n)
			l.freqs[freq+1] = nl
			l.keys[key] = n
		}
		return n.value
	} else {
		return -1
	}
}

func (l *LFUCache) Put(key int, value int) {
	if n, ok := l.keys[key]; ok {
		// 类似get
		freq := n.freq
		l.freqs[freq].Remove(n)
		if l.freqs[freq].size == 0 {
			delete(l.freqs, freq)
			if freq == l.minFreq {
				l.minFreq++
			}
		}
		n.freq++
		n.value = value
		if _, ok := l.freqs[freq+1]; ok {
			l.freqs[freq+1].Add(n)
		} else {
			nl := NewList2()
			nl.Add(n)
			l.freqs[freq+1] = nl
			l.keys[key] = n
		}
	} else {
		// 新建key->node, freq->list
		if len(l.keys) == l.capacity {
			// 已满需要移除minFreq
			f := l.freqs[l.minFreq]
			tail := f.tail.prev
			delete(l.keys, tail.key)
			f.Remove(tail)
			if l.freqs[l.minFreq].size == 0 {
				delete(l.freqs, l.minFreq)
			}
		}
		n := &Node2{
			key:   key,
			value: value,
			freq:  1,
		}
		l.keys[key] = n
		_, ok := l.freqs[1]
		if ok {
			l.freqs[1].Add(n)
		} else {
			nl := NewList2()
			nl.Add(n)
			l.freqs[1] = nl
		}
		l.minFreq = 1
	}
}

type List2 struct {
	head, tail *Node2
	size       int
}

func NewList2() *List2 {
	l := &List2{
		head: &Node2{},
		tail: &Node2{},
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

type Node2 struct {
	key, value, freq int
	prev, next       *Node2
}

func (l *List2) Remove(n *Node2) {
	n.prev.next = n.next
	n.next.prev = n.prev
	l.size--
}

func (l *List2) Add(n *Node2) {
	n.next = l.head.next
	n.prev = l.head
	l.head.next.prev = n
	l.head.next = n
	l.size++
}
