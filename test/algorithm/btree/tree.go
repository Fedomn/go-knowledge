package btree

import "errors"

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
func (tree *BPlusTree) Insert(key string, value string) error {
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

func (tree *BPlusTree) insertIntoParent(leftNode node, key string, value string, rightNode node) {
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

	spilledKey, spilledValue, spilledNode, spilled := leftNode.parent().insert(key, value, leftNode, rightNode)
	if !spilled {
		return
	}

	tree.insertIntoParent(leftNode, spilledKey, spilledValue, spilledNode)
}

func (tree BPlusTree) Search(key string) (string, bool) {
	ok, idx, foundInode := tree.search(tree.root, key)
	if !ok {
		return "", false
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

func (tree *BPlusTree) search(n node, key string) (exist bool, inodeIdx int, foundNode node) {
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
