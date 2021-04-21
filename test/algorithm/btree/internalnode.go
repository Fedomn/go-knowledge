package btree

import (
	"fmt"
	"sort"
)

// internal inode inside internalNode
type internalInode struct {
	key   int
	value int
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

func (in *internalNode) find(key int) (int, bool) {
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

func (in *internalNode) insert(key int, value int, leftNode node, rightNode node) (spilledKey int, spilledValue int, spilledNode node, spilled bool) {
	idx, _ := in.find(key)

	// Add capacity and shift nodes if we don't have an exact match and need to insert.
	exact := len(in.inodes) > 0 && idx < len(in.inodes)
	if !exact {
		in.inodes = append(in.inodes, internalInode{})
		copy(in.inodes[idx+1:], in.inodes[idx:])
	} else {
		in.inodes = append(in.inodes[:idx+1], in.inodes[idx:]...)
		// 插入中间的 internal node，需要更新next inode的left
		in.inodes[idx+1].left = rightNode
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
	return -1, -1, nil, false
}

func (in *internalNode) split() (node, int, int) {
	midIdx := in.degree / 2
	midKey := in.inodes[midIdx].key
	midValue := in.inodes[midIdx].value

	// new internalNode as rightBranch
	newInternalNode := newInternalNode(in.degree)
	newInternalNode.inodes = append(newInternalNode.inodes[0:], in.inodes[midIdx+1:]...)

	// update original internalNode
	updatedInodes := make([]internalInode, midIdx)
	copy(updatedInodes, in.inodes[:midIdx])
	in.inodes = updatedInodes

	return newInternalNode, midKey, midValue
}

func (in *internalNode) count() int {
	return len(in.inodes)
}

func (in *internalNode) largerOrEqualsThanHalfDegree() bool {
	return in.count() >= in.degree/2
}

func (in *internalNode) smallerThanHalfDegree() bool {
	return in.count() < in.degree/2
}

func (in *internalNode) delete(deletedKey int, deletedInodeIdx int) (freeRoot bool, newRoot node) {
	_, _, _, predecessorNode := in.predecessorMax(deletedInodeIdx)
	_, _, _, successorNode := in.successorMin(deletedInodeIdx)
	//if in.inodes[deletedInodeIdx].left.largerOrEqualsThanHalfDegree() {
	if predecessorNode.largerOrEqualsThanHalfDegree() {
		fmt.Println("borrowLeft")
		// 取前驱的最大值
		key, value, inodeIdx, leafNode := in.predecessorMax(deletedInodeIdx)
		in.inodes[deletedInodeIdx].key = key
		in.inodes[deletedInodeIdx].value = value
		// FIXME 先删除inode 再rebalance ?
		//leafNode.inodes = append(leafNode.inodes[:inodeIdx], leafNode.inodes[inodeIdx+1:]...)
		leafNode.delete(key, inodeIdx)
		return false, nil
		//} else if in.inodes[deletedInodeIdx].right.largerOrEqualsThanHalfDegree() {
	} else if successorNode.largerOrEqualsThanHalfDegree() {
		fmt.Println("borrowRight")
		// 取后继的最小值
		key, value, inodeIdx, leafNode := in.successorMin(deletedInodeIdx)
		in.inodes[deletedInodeIdx].key = key
		in.inodes[deletedInodeIdx].value = value
		// FIXME 先删除inode 再rebalance ?
		//leafNode.inodes = append(leafNode.inodes[:inodeIdx], leafNode.inodes[inodeIdx+1:]...)
		leafNode.delete(key, inodeIdx)
		return false, nil
	} else {
		// merge left and right branch
		// 将right branch copy到left branch，并将deletedKey的next key的left node指向 待merge的left branch
		rightBranch := in.inodes[deletedInodeIdx].right
		leftBranch := in.inodes[deletedInodeIdx].left
		switch right := rightBranch.(type) {
		case *leafNode:
			left := leftBranch.(*leafNode)
			left.inodes = append(left.inodes, right.inodes...)

			if deletedInodeIdx != len(in.inodes)-1 {
				in.inodes[deletedInodeIdx+1].left = left
				in.inodes = append(in.inodes[:deletedInodeIdx], in.inodes[deletedInodeIdx+1:]...)
			}
			// TODO need to handle more complex condition, now we not consider root and some corner condition
			return false, nil
		case *internalNode:
			left := leftBranch.(*internalNode)
			left.inodes = append(left.inodes, right.inodes...)

			if deletedInodeIdx != len(in.inodes)-1 {
				in.inodes[deletedInodeIdx+1].left = left
				in.inodes = append(in.inodes[:deletedInodeIdx], in.inodes[deletedInodeIdx+1:]...)
			}

			if in.parent() == nil {
				// reset tree root
				return true, left
			}

			if len(in.inodes) == 0 && in.parent() == nil {

			}

			return false, nil
			// TODO need to handle more complex condition, now we not consider root and some corner condition
		default:
			panic("invalid node")
		}
	}
}

func (in *internalNode) predecessorMax(deletedInodeIdx int) (key int, value int, inodeIdx int, node *leafNode) {
	// 当前node的前驱nodes里的最大值
	cursor := in.inodes[deletedInodeIdx].left
	for {
		switch node := cursor.(type) {
		case *leafNode:
			inode := node.inodes[node.count()-1]
			return inode.key, inode.value, node.count() - 1, node
		case *internalNode:
			cursor = node.inodes[node.count()-1].right
		default:
			panic("invalid node")
		}
	}
}

func (in *internalNode) successorMin(deletedInodeIdx int) (key int, value int, inodeIdx int, node *leafNode) {
	// 当前node的后继nodes里的最小值
	cursor := in.inodes[deletedInodeIdx].right
	for {
		switch node := cursor.(type) {
		case *leafNode:
			inode := node.inodes[0]
			return inode.key, inode.value, 0, node
		case *internalNode:
			cursor = node.inodes[0].left
		default:
			panic("invalid node")
		}
	}
}

func (in *internalNode) sibling(leafInodeKey int) (leftSibling, rightSibling node, leftSiblingIdx, rightSiblingIdx int) {
	firstLargerKeyIdx, _ := in.find(leafInodeKey)
	// left sub-tree keys < node-key <= right sub-tree keys

	if firstLargerKeyIdx == 0 {
		// first larger key idx = 0，所以只有右兄弟
		return nil, in.inodes[0].right, -1, 0
	} else if firstLargerKeyIdx == len(in.inodes) {
		// first larger key idx 不存在，所有只有左兄弟
		return in.inodes[len(in.inodes)-1].left, nil, len(in.inodes) - 1, -1
	} else {
		leftSibling = in.inodes[firstLargerKeyIdx-1].left
		rightSibling = in.inodes[firstLargerKeyIdx].right
		leftSiblingIdx = firstLargerKeyIdx - 1
		rightSiblingIdx = firstLargerKeyIdx
		return
	}
}
