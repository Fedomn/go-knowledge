package btree

import (
	"errors"
)

type BPlusTree struct {
	root   node
	degree int
}

func newTree(degree int) *BPlusTree {
	return &BPlusTree{
		degree: degree,
	}
}

// key points
// 1. 通过tree的insert最终一定是发生在leafNode上，internalNode上insert只会发生在spill时刻需要上溢节点
func (tree *BPlusTree) Insert(key int, value int) error {
	if tree.root == nil {
		leafNode := newLeafNode(tree.degree)
		leafNode.insert(key, value, nil, nil)
		tree.root = leafNode
		return nil
	}

	// 找到待插入的leafNode
	ok, _, foundNode := tree.search(tree.root, key)
	if ok {
		return errors.New("already exist same key")
	}

	spilledKey, spilledValue, spilledNode, spilled := foundNode.insert(key, value, nil, nil)
	if !spilled {
		return nil
	}

	tree.insertIntoParent(foundNode, spilledKey, spilledValue, spilledNode)
	return nil
}

func (tree *BPlusTree) insertIntoParent(leftNode node, key int, value int, rightNode node) {
	if leftNode.parent() == nil {
		internalNode := newInternalNode(tree.degree)
		internalNode.insert(key, value, leftNode, rightNode)

		if tree.root == leftNode {
			tree.root = internalNode
		}
		leftNode.setParent(internalNode)
		rightNode.setParent(internalNode)
		return
	}

	leftNodeParent := leftNode.parent()

	pSpilledKey, pSpilledValue, pSpilledNode, pSpilled := leftNodeParent.insert(key, value, leftNode, rightNode)
	if !pSpilled {
		return
	}

	// 递归到这里的情况，一定是internalNode的分裂，所以这时insertIntoParent的leftNode是leftNodeParent
	// 同时递归之前，将internalNode对应的左右leafNode的parent值给更新好(只要更新pSpilledNode的leafNode就好，因为左branch仍用老的node)
	if pNode, ok := pSpilledNode.(*internalNode); ok {
		for idx := range pNode.inodes {
			pNode.inodes[idx].left.setParent(pNode)
			pNode.inodes[idx].right.setParent(pNode)
		}
	}
	tree.insertIntoParent(leftNodeParent, pSpilledKey, pSpilledValue, pSpilledNode)
}

func (tree BPlusTree) Search(key int) (int, bool) {
	ok, idx, foundInode := tree.search(tree.root, key)
	if !ok {
		return -1, false
	}
	switch node := foundInode.(type) {
	case *leafNode:
		return node.inodes[idx].value, true
	case *internalNode:
		return node.inodes[idx].value, true
	default:
		panic("invalid node")
	}
}

func (tree *BPlusTree) search(n node, key int) (exist bool, inodeIdx int, foundNode node) {
	cursor := n
	inodeIdx = -1
	for {
		switch node := cursor.(type) {
		case *leafNode:
			idx, ok := node.find(key)
			return ok, idx, node
		case *internalNode:
			// found inodeIdx associated key always >= given key
			idx, ok := node.find(key)
			if ok {
				return ok, idx, node
			}

			// handle not found key that inodeIdx = len(node.inodes)
			if idx >= len(node.inodes) {
				idx = idx - 1
			}

			if key < node.inodes[idx].key {
				cursor = node.inodes[idx].left
			} else if key > node.inodes[idx].key {
				cursor = node.inodes[idx].right
			}

			inodeIdx = idx
		default:
			panic("invalid node")
		}
	}
}

func (tree *BPlusTree) preOrderTraversal() []int {
	result := make([]int, 0)
	cursor := tree.root
	preOrderWalk(cursor, &result)
	return result
}

func preOrderWalk(cursor node, result *[]int) {
	switch node := cursor.(type) {
	case *leafNode:
		for _, inode := range node.inodes {
			*result = append(*result, inode.key)
		}
	case *internalNode:
		for i := range node.inodes {
			inode := node.inodes[i]
			*result = append(*result, inode.key)
			preOrderWalk(inode.left, result)
			preOrderWalk(inode.right, result)
		}
	default:
		panic("invalid node")
	}
}
