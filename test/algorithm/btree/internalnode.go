package btree

import (
	"sort"
)

// internal inode inside internalNode
type internalInode struct {
	key   string
	left  node
	right node
}

type internalNode struct {
	p      *internalNode
	inodes []internalInode
}

func newInternalNode(p *internalNode) *internalNode {
	in := &internalNode{
		p:      p,
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

// max degree in internal node, it depends on os pageSize - internalNode pointer size / internalInode size
const MaxInternalInodesCount = 3

func (in *internalNode) full() bool {
	return len(in.inodes) == MaxInternalInodesCount
}

func (in *internalNode) insert(key string, child node) (spilledKey string, spilledNode *internalNode, spilled bool) {
	idx, _ := in.find(key)

	in.inodes = append(in.inodes[idx+1:], in.inodes[idx:len(in.inodes)]...)
	in.inodes[idx].key = key
	// TODO need fill inode left and right keys
	//in.inodes[idx]. = child
	//child.setParent(in)

	if in.full() {
		rightBranchNode, midKey := in.split()
		return midKey, rightBranchNode, true
	}
	return "", nil, false
}

func (in *internalNode) split() (*internalNode, string) {
	midIdx := MaxInternalInodesCount / 2
	midKey := in.inodes[midIdx].key

	// new internalNode as rightBranch
	newInternalNode := newInternalNode(nil)
	newInternalNode.inodes = append(newInternalNode.inodes[0:], in.inodes[midIdx:]...)

	// update new internal node children's parent
	//for i := 0; i < len(newInternalNode.inodes); i++ {
		//newInternalNode.inodes[i].child.setParent(newInternalNode)
		//left := newInternalNode.inodes[i].left
	//}

	// update original internalNode
	updatedInodes := make([]internalInode, midIdx)
	copy(updatedInodes, in.inodes[:midIdx])
	in.inodes = updatedInodes

	return newInternalNode, midKey
}
