package btree

import (
	"sort"
)

// leaf inode inside leafNode
type leafInode struct {
	key   string
	value string
}

type leafNode struct {
	degree int
	inodes []leafInode
}

func newLeafNode(degree int) *leafNode {
	return &leafNode{
		degree: degree,
	}
}

func (l *leafNode) find(key string) (int, bool) {
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

//func (l *leafNode) parent() *internalNode {
//	return l.p
//}
//
//func (l *leafNode) setParent(in *internalNode) {
//	l.p = in
//}

// max degree in internal node, it depends on os pageSize - internalNode pointer size / internalInode size
const MaxLeafInodesCount = 3

func (l *leafNode) full() bool {
	return len(l.inodes) == MaxLeafInodesCount
}

func (l *leafNode) insert(key string, value string) (spilledKey string, spilledNode node, spilled bool) {
	idx, _ := l.find(key)
	l.inodes = append(l.inodes[idx+1:], l.inodes[idx:]...)
	l.inodes[idx].key = key
	l.inodes[idx].key = value

	if l.full() {
		nextLeafNode := l.split()
		return nextLeafNode.inodes[0].key, nextLeafNode, true
	}

	return "", nil, false
}

func (l *leafNode) split() *leafNode {
	newLeafNode := newLeafNode(nil)
	midIdx := MaxInternalInodesCount / 2

	newLeafNode.inodes = make([]leafInode, len(l.inodes)-midIdx)
	copy(newLeafNode.inodes, l.inodes[midIdx:])
	return newLeafNode
}
