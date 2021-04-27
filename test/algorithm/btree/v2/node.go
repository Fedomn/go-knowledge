package v2

import (
	"fmt"
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

// maximum count = (t-1) left part + (1) spilled key + (t-1) right part
func (n *node) isFull() bool {
	return n.count == n.t*2-1
}

// 理论上一个node最少要有t-1 keys，但是这里的判断是 >= t，原因是:
// 1. 用于removeFromNonLeaf函数时 只有 >=t，才可以删除，因为有最少t-1限制
// 2. 用于fill函数时，只有 >=t，才能在borrow 1个后，仍然满足最少t-1限制
// 3. 用于remove函数逐层搜索中，如果不满足 >=t，就需要先fill，否则待删除key后，节点就不满足t-1限制
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

func (n *node) remove(key int, b *btree) {
	idx, found := n.find(key)
	if found {
		if n.isLeaf {
			n.removeFromLeaf(idx)
		} else {
			n.removeFromNonLeaf(idx)
		}
	} else {
		// 由于removeFromNonLeaf里递归地调用remove
		// 因此会从上至下 逐层寻找 需要删除的predKey或succKey
		// 在这个过程中，填充那些key个数不足的node

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
			n.children[idx-1].remove(key, b)
		} else {
			n.children[idx].remove(key, b)
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

func (n *node) removeFromNonLeaf(idx int) {
	fmt.Printf("removeFromNonLeaf %d\n", idx)
	removedKey := n.keys[idx]

	if n.children[idx].overMiniDegree() {
		predKey := n.getPred(idx)
		n.keys[idx] = predKey
		// recursively delete pred key in leftChild
		n.children[idx].remove(predKey, nil)
	} else if n.children[idx+1].overMiniDegree() {
		succKey := n.getSucc(idx)
		n.keys[idx] = succKey
		// recursively delete succ key in rightChild
		n.children[idx+1].remove(succKey, nil)
	} else {
		// merge right child to left child and then delete key in left child
		n.merge(idx)
		n.children[idx].remove(removedKey, nil)
	}
}

// get predecessor maximum key
func (n *node) getPred(idx int) int {
	fmt.Printf("getPred %d\n", idx)
	leftChild := n.children[idx]

	for !leftChild.isLeaf {
		leftChild = leftChild.children[leftChild.count]
	}

	return leftChild.keys[leftChild.count-1]
}

// get successor minimum key
func (n *node) getSucc(idx int) int {
	fmt.Printf("getSucc %d\n", idx)
	rightChild := n.children[idx+1]

	for !rightChild.isLeaf {
		rightChild = rightChild.children[0]
	}

	return rightChild.keys[0]
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

	// underflow parent key into insertedKey, and then up borrowed key into parent
	insertedChild.keys[0] = n.keys[idx-1]
	if !insertedChild.isLeaf {
		insertedChild.children[0] = leftSiblingChild.children[leftSiblingChild.count]
	}

	n.keys[idx-1] = leftSiblingChild.keys[leftSiblingChild.count-1]
	leftSiblingChild.keys[leftSiblingChild.count-1] = -1

	insertedChild.count++
	leftSiblingChild.count--
}

func (n *node) borrowFromNext(idx int) {
	fmt.Printf("borrowFromNext %d\n", idx)
	insertedChild := n.children[idx]
	rightSiblingChild := n.children[idx+1]

	insertedChild.keys[insertedChild.count] = n.keys[idx]
	if !insertedChild.isLeaf {
		insertedChild.children[insertedChild.count+1] = rightSiblingChild.children[0]
	}

	n.keys[idx] = rightSiblingChild.keys[0]

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

// merge right child into left child, idx is left child index
// precondition: left and right child count < miniDegree
func (n *node) merge(idx int) {
	fmt.Printf("merge %d\n", idx)
	leftChild := n.children[idx]
	rightChild := n.children[idx+1]

	// removedKey underflow into leftChild
	leftChild.keys[leftChild.count] = n.keys[idx]

	// copy right child to left child
	for i := 0; i < rightChild.count; i++ {
		leftChild.keys[i+leftChild.count+1] = rightChild.keys[i]
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
	leftChild.count += rightChild.count + 1
	n.count--

	rightChild = nil
}
