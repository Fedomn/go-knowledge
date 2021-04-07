package btree

import (
	"sort"
)

// internal inode inside internalNode
type internalInode struct {
	key   string
	value string
	left  node
	right node
}

type internalNode struct {
	degree int
	p      *internalNode
	inodes []internalInode
}

func newInternalNode(degree int) *internalNode {
	in := &internalNode{
		degree: degree,
		inodes: make([]internalInode, 0),
	}
	return in
}

func (in *internalNode) find(key string) (int, bool) {
	comparedFunc := func(idx int) bool {
		return in.inodes[idx].key >= key
	}

	i := sort.Search(len(in.inodes), comparedFunc)

	if i < len(in.inodes) && in.inodes[i].key == key {
		return i, true
	}

	return i, false
}

func (in *internalNode) parent() *internalNode {
	return in.p
}

func (in *internalNode) setParent(i *internalNode) {
	in.p = i
}

func (in *internalNode) full() bool {
	return len(in.inodes) == in.degree
}

func (in *internalNode) insert(key string, value string, leftNode node, rightNode node) (spilledKey string, spilledValue string, spilledNode node, spilled bool) {
	idx, _ := in.find(key)

	// Add capacity and shift nodes if we don't have an exact match and need to insert.
	exact := len(in.inodes) > 0 && idx < len(in.inodes)
	if !exact {
		in.inodes = append(in.inodes, internalInode{})
		copy(in.inodes[idx+1:], in.inodes[idx:])
	} else {
		in.inodes = append(in.inodes[idx+1:], in.inodes[idx:]...)
		// in.inodes = append(in.inodes[idx+1:], in.inodes[idx:len(in.inodes)]...)
	}

	in.inodes[idx].key = key
	in.inodes[idx].value = value
	in.inodes[idx].left = leftNode
	in.inodes[idx].left.setParent(in)
	in.inodes[idx].right = rightNode
	in.inodes[idx].right.setParent(in)

	if in.full() {
		rightBranchNode, midKey, midValue := in.split()
		return midKey, midValue, rightBranchNode, true
	}
	return "", "", nil, false
}

func (in *internalNode) split() (node, string, string) {
	midIdx := in.degree / 2
	midKey := in.inodes[midIdx].key
	midValue := in.inodes[midIdx].value

	// new internalNode as rightBranch
	newInternalNode := newInternalNode(in.degree)
	// TODO
	newInternalNode.inodes = append(newInternalNode.inodes[0:], in.inodes[midIdx+1:]...)

	// update new internal node children's parent
	//for i := 0; i < len(newInternalNode.inodes); i++ {
	//newInternalNode.inodes[i].child.setParent(newInternalNode)
	//left := newInternalNode.inodes[i].left
	//}

	// update original internalNode
	updatedInodes := make([]internalInode, midIdx)
	copy(updatedInodes, in.inodes[:midIdx])
	in.inodes = updatedInodes

	return newInternalNode, midKey, midValue
}
