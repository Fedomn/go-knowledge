package bplustree

import (
	"fmt"
	"sort"
)

type node struct {
	isLeaf   bool
	t        int     // minimum degree (defines the range for number of keys)
	keys     []int   // inserted keys
	count    int     // current number of keys
	children []*node // child pointers
	// changed for bplustree
	next *node // next sibling
}

func newNode(miniDegree int, isLeaf bool) *node {
	n := &node{
		isLeaf:   isLeaf,
		t:        miniDegree,
		keys:     make([]int, 2*miniDegree-1),
		count:    0,
		children: make([]*node, 2*miniDegree),
		next:     nil,
	}
	for i := range n.keys {
		n.keys[i] = -1
	}
	return n
}

// maximum count = (t-1) left part + (1) spilled key + (t-1) right part
func (n *node) isFull() bool {
	return n.count == n.t*2-1
}

// 理论上一个node最少要有t-1 keys，但是这里的判断是 >= t，原因是:
// 1. 用于fill函数时，只有 >=t，才能在borrow 1个后，仍然满足最少t-1限制
// 2. 用于remove函数逐层搜索中，如果不满足 >=t，就需要先fill，否则待删除key后，节点就不满足t-1限制
func (n *node) overMiniDegree() bool {
	return n.count >= n.t
}

func (n *node) hasDuplicatedKey(key int) bool {
	for _, k := range n.keys {
		if key == k {
			return true
		}
	}
	return false
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

// changed for bplustree
func (n *node) firstKey() int {
	if n.isLeaf {
		return n.keys[0]
	}

	return n.children[0].firstKey()
}

func (n *node) insertNonFull(key int) {
	// index start rightmost element
	idx := n.count - 1

	if n.isLeaf {
		// ignore duplicated key
		if n.hasDuplicatedKey(key) {
			return
		}

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

		// ignore duplicated key
		if n.hasDuplicatedKey(key) || n.children[idx+1].hasDuplicatedKey(key) {
			return
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

// changed for bplustree
// split oldNode to two children nodes and copied spilledKey to n which as parent node
// so n is parent node, oldNode as n's left child, inner function will create right node
// oldNode must be full (contains 2*t-1 keys)
// if oldNode is leaf node, then move t keys to rightNode
// if oldNode is non-leaf node, then move t-1 keys to rightNode
func (n *node) splitChild(oldNodeInParentIdx int, oldNode *node) {
	rightNode := newNode(oldNode.t, oldNode.isLeaf)
	if oldNode.isLeaf {
		rightNode.count = n.t
		// move last t keys of oldNode to rightNode
		for i := 0; i < n.t; i++ {
			rightNode.keys[i] = oldNode.keys[i+n.t-1]
		}
	} else {
		rightNode.count = n.t - 1
		// move last t-1 keys of oldNode to rightNode
		for i := 0; i < n.t-1; i++ {
			rightNode.keys[i] = oldNode.keys[i+n.t]
		}
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

	if oldNode.isLeaf {
		n.keys[oldNodeInParentIdx] = oldNode.keys[n.t-1]
		// reset oldNode keys
		for i := n.t - 1; i < oldNode.count; i++ {
			oldNode.keys[i] = -1
		}
	} else {
		n.keys[oldNodeInParentIdx] = oldNode.keys[n.t-1]
		// reset oldNode keys
		for i := n.t - 1; i < oldNode.count; i++ {
			oldNode.keys[i] = -1
		}
	}

	oldNode.count = n.t - 1
	n.count = n.count + 1

	// changed for bplustree
	if oldNode.isLeaf {
		rightNode.next = oldNode.next
		oldNode.next = rightNode
	}
}

func (n *node) remove(key int) {
	idx, found := n.find(key)
	if found && n.isLeaf {
		n.removeFromLeaf(idx)
	} else {
		// 递归往下调用remove
		// 在这个过程中，填充那些key个数不足的node

		// 如果key出现在internal node中，则idx为first largerOrEquals位置，也就是说n.keys[idx] = key
		// 所以，key真实存在的child是idx+1
		if found {
			idx++
		}

		if n.isLeaf {
			fmt.Printf("The key %d dose not exist in the tree\n", key)
			//fmt.Printf("%v\n", b.breadthFirstDraw())
			return
		}

		// flag indicate whether the key is the last one or not
		flag := idx == n.count

		// idx 是first LargerOrEquals key的位置，所以需要插入的是left child，即children[idx]
		if !n.children[idx].overMiniDegree() {
			n.fill(idx)
		}

		// 继续往下搜索key
		// flag表明之前是last child，idx>n.count表明执行了上面的fill操作里的merge
		// 所以，导致n上的key underflow，所以left child的index 应该向左移一位，即idx-1
		if flag && idx > n.count {
			n.children[idx-1].remove(key)
		} else {
			n.children[idx].remove(key)
		}
	}
}

func (n *node) removeFromLeaf(idx int) {
	fmt.Printf("removeFromLeaf %d\n", idx)
	for i := idx + 1; i < n.count; i++ {
		n.keys[i-1] = n.keys[i]
	}
	n.keys[n.count-1] = -1
	n.count--
}

func (n *node) fill(idx int) {
	fmt.Printf("fill %d\n", idx)
	if idx != 0 && n.children[idx-1].overMiniDegree() {
		n.borrowFromPrev(idx)
	} else if idx != n.count && n.children[idx+1].overMiniDegree() {
		n.borrowFromNext(idx)
	} else {
		if idx != n.count {
			n.merge(idx)
		} else {
			n.merge(idx - 1)
		}
	}
}

// changed for bplustree
func (n *node) borrowFromPrev(idx int) {
	fmt.Printf("borrowFromPrev %d\n", idx)
	insertedChild := n.children[idx]
	leftSiblingChild := n.children[idx-1]

	// prepare a new slot for borrowed key from left sibling
	for i := insertedChild.count - 1; i >= 0; i-- {
		insertedChild.keys[i+1] = insertedChild.keys[i]
	}
	if !insertedChild.isLeaf {
		// move child pointer if needed
		for i := insertedChild.count; i >= 0; i-- {
			insertedChild.children[i+1] = insertedChild.children[i]
		}
	}

	underflowKey := n.keys[idx-1]
	if insertedChild.isLeaf {
		// use leftSiblingChild maximum key to replace underflow key in bplustree,
		// because bplustree rightChild keys >= spilledKey
		insertedChild.keys[0] = leftSiblingChild.keys[leftSiblingChild.count-1]
	} else {
		// underflow parent key into insertedKey, and then up borrowed key into parent
		insertedChild.keys[0] = underflowKey
		insertedChild.children[0] = leftSiblingChild.children[leftSiblingChild.count]
	}

	n.keys[idx-1] = leftSiblingChild.keys[leftSiblingChild.count-1]
	leftSiblingChild.keys[leftSiblingChild.count-1] = -1

	insertedChild.count++
	leftSiblingChild.count--
}

// changed for bplustree
func (n *node) borrowFromNext(idx int) {
	fmt.Printf("borrowFromNext %d\n", idx)
	insertedChild := n.children[idx]
	rightSiblingChild := n.children[idx+1]

	// 注意，B+tree里 n.keys[idx] 一定是等于 rightSiblingChild.keys[0]
	if insertedChild.isLeaf {
		insertedChild.keys[insertedChild.count] = rightSiblingChild.keys[0]
		// pick second one
		n.keys[idx] = rightSiblingChild.keys[1]
	} else {
		insertedChild.keys[insertedChild.count] = n.keys[idx]
		insertedChild.children[insertedChild.count+1] = rightSiblingChild.children[0]
		n.keys[idx] = rightSiblingChild.keys[0]
	}

	for i := 1; i < rightSiblingChild.count; i++ {
		rightSiblingChild.keys[i-1] = rightSiblingChild.keys[i]
	}
	// reset borrowed key position
	rightSiblingChild.keys[rightSiblingChild.count-1] = -1

	if !insertedChild.isLeaf {
		for i := 1; i <= rightSiblingChild.count; i++ {
			rightSiblingChild.children[i-1] = rightSiblingChild.children[i]
		}
	}

	insertedChild.count++
	rightSiblingChild.count--
}

// changed for bplustree
// merge right child into left child, idx is left child index
// precondition: left and right child count < miniDegree
func (n *node) merge(idx int) {
	fmt.Printf("merge %d\n", idx)
	leftChild := n.children[idx]
	rightChild := n.children[idx+1]

	// copy right child to left child
	if leftChild.isLeaf {
		for i := 0; i < rightChild.count; i++ {
			leftChild.keys[i+leftChild.count] = rightChild.keys[i]
		}
	} else {
		// removedKey underflow into leftChild
		leftChild.keys[leftChild.count] = n.keys[idx]

		for i := 0; i < rightChild.count; i++ {
			leftChild.keys[i+leftChild.count+1] = rightChild.keys[i]
		}
	}

	// copy right child pointers to left child
	if !leftChild.isLeaf {
		for i := 0; i <= rightChild.count; i++ {
			leftChild.children[i+leftChild.count+1] = rightChild.children[i]
		}
	}

	// since removedKey underflow, so fill the gap by moving forward keys
	for i := idx + 1; i < n.count; i++ {
		n.keys[i-1] = n.keys[i]
	}
	// fill child pointer gap
	for i := idx + 2; i <= n.count; i++ {
		n.children[i-1] = n.children[i]
	}

	// update count
	if leftChild.isLeaf {
		leftChild.count += rightChild.count
	} else {
		leftChild.count += rightChild.count + 1
	}

	// update next pointer
	if leftChild.isLeaf {
		leftChild.next = rightChild.next
	}

	n.count--
	rightChild = nil
}
