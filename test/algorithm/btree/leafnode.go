package btree

import (
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
