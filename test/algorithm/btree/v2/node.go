package v2

import (
	"sort"
)

// keys and children index relationship:
// when keys idx = 0
// its left child is child idx 0
// its right child is child idx 1

type node struct {
	isLeaf   bool
	t        int     // minimum degree (defines the range for number of keys)
	keys     []int   // inserted keys
	count    int     // current number of keys
	children []*node // child pointers
}

func newNode(miniDegree int, isLeaf bool) *node {
	n := &node{
		isLeaf:   isLeaf,
		t:        miniDegree,
		keys:     make([]int, 2*miniDegree-1),
		count:    0,
		children: make([]*node, 2*miniDegree),
	}
	for i := range n.keys {
		n.keys[i] = -1
	}
	return n
}

func (n *node) isFull() bool {
	return n.count == n.t*2-1
}

// find the first key greater than or equal to k
// return 0 when k <= first smallest key
// return count when k > last largest key
// otherwise, return idx that in range [smallest key idx, largest key idx]
func (n *node) find(key int) (int, bool) {
	idx := sort.Search(n.count, func(i int) bool {
		return n.keys[i] >= key
	})

	if idx < n.count && key == n.keys[idx] {
		return idx, true
	}
	return idx, false
}

// search a key in subtree rooted with this node, not found return nil
func (n *node) search(key int) (*node, bool) {
	// idx is first largerOrEquals than key
	idx, found := n.find(key)
	if found {
		return n, true
	}

	if n.isLeaf {
		return nil, false
	}

	// search left child
	return n.children[idx].search(key)
}

// split oldNode to two children nodes and moved spilledKey to n which as parent node
// so n is parent node, oldNode as n's left child, inner function will create right node
// oldNode must be full (contains 2*t-1 keys)
// move t-1 keys to rightNode
func (n *node) splitChild(oldNodeInParentIdx int, oldNode *node) {
	rightNode := newNode(oldNode.t, oldNode.isLeaf)
	rightNode.count = n.t - 1

	// move last t-1 keys of oldNode to rightNode
	for i := 0; i < n.t-1; i++ {
		rightNode.keys[i] = oldNode.keys[i+n.t]
	}

	// for internalNode move last t child pointers of oldNode to rightNode
	if !oldNode.isLeaf {
		for i := 0; i < n.t; i++ {
			rightNode.children[i] = oldNode.children[i+n.t]
		}
	}

	// since this node will have a new child, so create space for new child
	// oldNodeInParentIdx is need inserted position
	// 从oldNodeInParentIdx+1开始整体向后移动一位，为新来的child准备位置
	for i := n.count; i >= oldNodeInParentIdx+1; i-- {
		n.children[i+1] = n.children[i]
	}

	// update child pointer
	n.children[oldNodeInParentIdx+1] = rightNode

	// A key of y will move to this node. Find location of
	// new key and move all greater keys one space ahead
	for i := n.count - 1; i >= oldNodeInParentIdx; i-- {
		n.keys[i+1] = n.keys[i]
	}

	n.keys[oldNodeInParentIdx] = oldNode.keys[n.t-1]
	// reset oldNode keys
	for i := n.t - 1; i < oldNode.count; i++ {
		oldNode.keys[i] = -1
	}
	// reset oldNode count
	oldNode.count = n.t - 1

	n.count = n.count + 1
}

func (n *node) insertNonFull(key int) {
	// index start rightmost element
	idx := n.count - 1

	if n.isLeaf {
		// from rightmost to left, find the location of new key to be inserted
		// move all greater keys to right as finding loop
		for idx >= 0 && n.keys[idx] > key {
			n.keys[idx+1] = n.keys[idx]
			idx--
		}

		n.keys[idx+1] = key
		n.count++
	} else {
		// find the child which will have the new key
		for idx >= 0 && n.keys[idx] > key {
			idx--
		}

		// idx is the first larger index than key, so will insert to rightNode
		// see if the found child is full
		if n.children[idx+1].isFull() {
			n.splitChild(idx+1, n.children[idx+1])

			if n.keys[idx+1] < key {
				idx++
			}
		}
		n.children[idx+1].insertNonFull(key)
	}
}
