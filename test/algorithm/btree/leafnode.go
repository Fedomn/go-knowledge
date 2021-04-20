package btree

import (
	"fmt"
	"sort"
)

// leaf inode inside leafNode
type leafInode struct {
	key   int
	value int
}

type leafNode struct {
	degree int
	p      *internalNode
	inodes []leafInode
}

func newLeafNode(degree int) *leafNode {
	return &leafNode{
		degree: degree,
		inodes: make([]leafInode, 0),
	}
}

func (l *leafNode) find(key int) (int, bool) {
	comparedFunc := func(idx int) bool {
		// >= means find the first larger or equals key
		return l.inodes[idx].key >= key
	}

	i := sort.Search(len(l.inodes), comparedFunc)

	if i < len(l.inodes) && l.inodes[i].key == key {
		return i, true
	}
	return i, false
}

func (l *leafNode) parent() *internalNode {
	return l.p
}

func (l *leafNode) setParent(in *internalNode) {
	l.p = in
}

func (l *leafNode) full() bool {
	return len(l.inodes) == l.degree
}

func (l *leafNode) insert(key int, value int, leftNode node, rightNode node) (spilledKey int, spilledValue int, spilledNode node, spilled bool) {
	idx, _ := l.find(key)

	// Add capacity and shift nodes if we don't have an exact match and need to insert.
	exact := len(l.inodes) > 0 && idx < len(l.inodes)
	if !exact {
		l.inodes = append(l.inodes, leafInode{})
		copy(l.inodes[idx+1:], l.inodes[idx:])
	} else {
		// slice is "[ )" pattern
		l.inodes = append(l.inodes[:idx+1], l.inodes[idx:]...)
	}

	l.inodes[idx].key = key
	l.inodes[idx].value = value

	if l.full() {
		nextLeafNode, midKey, midValue := l.split()
		return midKey, midValue, nextLeafNode, true
	}

	return -1, -1, nil, false
}

func (l *leafNode) split() (node, int, int) {
	midIdx := l.degree / 2
	midKey := l.inodes[midIdx].key
	midValue := l.inodes[midIdx].value

	newLeafNode := newLeafNode(l.degree)
	newLeafNode.inodes = make([]leafInode, len(l.inodes)-midIdx-1)
	copy(newLeafNode.inodes, l.inodes[midIdx+1:])

	updatedInodes := make([]leafInode, midIdx)
	copy(updatedInodes, l.inodes[:midIdx])
	l.inodes = updatedInodes

	return newLeafNode, midKey, midValue
}

func (l *leafNode) count() int {
	return len(l.inodes)
}

func (l *leafNode) largerOrEqualsThanHalfDegree() bool {
	return l.count() >= l.degree/2
}

func (l *leafNode) smallerThanHalfDegree() bool {
	return l.count() < l.degree/2
}

func (l *leafNode) delete(key int, inodeIdx int) {
	l.inodes = append(l.inodes[:inodeIdx], l.inodes[inodeIdx+1:]...)
	if l.smallerThanHalfDegree() {
		// 1. borrow from left sibling node or right sibling node
		leftSibling, rightSibling, leftSiblingIdx, rightSiblingIdx := l.parent().sibling(key)

		if leftSibling != nil && leftSibling.largerOrEqualsThanHalfDegree() {
			l.borrowFromLeftSibling(leftSibling, leftSiblingIdx)
		} else if rightSibling != nil && rightSibling.largerOrEqualsThanHalfDegree() {
			l.borrowFromRightSibling(rightSibling, rightSiblingIdx)
		} else {
			// 2. If both the immediate sibling nodes already have a minimum number of keys,
			// then merge the node with either the left sibling node or the right sibling node.
			if leftSibling != nil {
				l.mergeLeftSibling(leftSibling, leftSiblingIdx)
			} else if rightSibling != nil {
				l.mergeRightSibling(rightSibling, rightSiblingIdx)
			}
		}
	}
}

func (l *leafNode) borrowFromLeftSibling(leftSibling node, leftSiblingIdx int) {
	fmt.Println("borrowFromLeftSibling")
	parentNode := l.parent()
	inode := parentNode.inodes[leftSiblingIdx]
	l.insert(inode.key, inode.value, nil, nil)

	// rotate leftSiblingNode last inode to parent inode
	leftSiblingNode := leftSibling.(*leafNode)
	deletedLeafInode := leftSiblingNode.inodes[leftSiblingNode.count()-1]
	parentNode.inodes[leftSiblingIdx].key = deletedLeafInode.key
	parentNode.inodes[leftSiblingIdx].value = deletedLeafInode.value

	// delete leftSiblingNode last inode
	leftSiblingNode.inodes = leftSiblingNode.inodes[:leftSiblingNode.count()-1]
}

func (l *leafNode) borrowFromRightSibling(rightSibling node, rightSiblingIdx int) {
	fmt.Println("borrowFromRightSibling")
	parentNode := l.parent()
	inode := parentNode.inodes[rightSiblingIdx]
	l.insert(inode.key, inode.value, nil, nil)

	// rotate rightSiblingNode first inode to parent inode
	rightSiblingNode := rightSibling.(*leafNode)
	deletedLeafInode := rightSiblingNode.inodes[0]
	parentNode.inodes[rightSiblingIdx].key = deletedLeafInode.key
	parentNode.inodes[rightSiblingIdx].value = deletedLeafInode.value

	// delete rightSiblingNode first inode
	rightSiblingNode.inodes = rightSiblingNode.inodes[1:]
}

func (l *leafNode) mergeLeftSibling(leftSibling node, leftSiblingIdx int) {
	fmt.Println("mergeLeftSibling")

	parentNode := l.parent()
	inode := parentNode.inodes[leftSiblingIdx]

	parentNode.inodes[leftSiblingIdx+1].left = leftSibling

	leftSibling.insert(inode.key, inode.value, nil, nil)
	for _, leafInode := range l.inodes {
		leftSibling.insert(leafInode.key, leafInode.value, nil, nil)
	}
	parentNode.inodes = append(parentNode.inodes[:leftSiblingIdx], parentNode.inodes[leftSiblingIdx+1:]...)
}

func (l *leafNode) mergeRightSibling(rightSibling node, rightSiblingIdx int) {
	fmt.Println("mergeRightSibling")

	parentNode := l.parent()
	inode := parentNode.inodes[rightSiblingIdx]

	parentNode.inodes[rightSiblingIdx+1].left = rightSibling

	rightSibling.insert(inode.key, inode.value, nil, nil)
	for _, leafInode := range l.inodes {
		rightSibling.insert(leafInode.key, leafInode.value, nil, nil)
	}
	parentNode.inodes = append(parentNode.inodes[:rightSiblingIdx], parentNode.inodes[rightSiblingIdx+1:]...)
}
