package btree

import (
	"strconv"
	"strings"
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

// Insert key points
// 1. 通过tree的insert最终一定是发生在leafNode上，internalNode上insert只会发生在spill时刻需要上溢节点
func (tree *BPlusTree) Insert(key int, value int) {
	if tree.root == nil {
		leafNode := newLeafNode(tree.degree)
		leafNode.insert(key, value, nil, nil)
		tree.root = leafNode
		return
	}

	// notFound -> 找到待插入的leafNode
	// found -> 找到待更新的leafNode / internalNode
	ok, inodeIdx, foundNode := tree.search(tree.root, key)
	if ok {
		switch node := foundNode.(type) {
		case *leafNode:
			node.inodes[inodeIdx].value = value
		case *internalNode:
			node.inodes[inodeIdx].value = value
		default:
			panic("invalid node")
		}
		return
	}

	spilledKey, spilledValue, spilledNode, spilled := foundNode.insert(key, value, nil, nil)
	if !spilled {
		return
	}

	tree.insertIntoParent(foundNode, spilledKey, spilledValue, spilledNode)
	return
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

func (tree *BPlusTree) Delete(key int) {
	ok, inodeIdx, foundNode := tree.search(tree.root, key)
	if !ok {
		return
	}

	switch node := foundNode.(type) {
	case *leafNode:
		node.delete(key, inodeIdx)
	case *internalNode:
		if freeRoot, newRoot := node.delete(key, inodeIdx); freeRoot {
			tree.root = newRoot
		}
	default:
		panic("invalid node")
	}
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

// for draw
// use tools: https://projects.calebevans.me/b-sketcher/
func (tree *BPlusTree) breadthFirstWalk() []node {
	bfsResult := make([]node, 0)
	stack := make([]node, 0)
	nextLayerStack := make([]node, 0)
	stack = append(stack, tree.root)
	for len(stack) > 0 {
		cursor := stack[0]
		// into bfsResult
		bfsResult = append(bfsResult, cursor)
		if len(stack) >= 1 {
			stack = stack[1:]
		} else {
			stack = []node{}
		}

		switch node := cursor.(type) {
		case *leafNode:
			// impossible to happen
			continue
		case *internalNode:
			for i := 0; i < len(node.inodes); i++ {
				inode := node.inodes[i]
				nextLayerStack = append(nextLayerStack, inode.left)
				if i == len(node.inodes)-1 {
					nextLayerStack = append(nextLayerStack, inode.right)
				}
			}
		default:
			panic("invalid node")
		}

		// stack为空，代表这一层出栈完了，可以添加分隔符了
		if len(stack) == 0 {
			bfsResult = append(bfsResult, nil)
			stack = nextLayerStack
			nextLayerStack = make([]node, 0)
		}
	}
	return bfsResult
}

func (tree *BPlusTree) breadthFirstDraw() string {
	var result = ""
	walkNodes := tree.breadthFirstWalk()
	for _, cursor := range walkNodes {
		switch node := cursor.(type) {
		case *leafNode:
			for idx, inode := range node.inodes {
				result += strconv.Itoa(inode.key)
				if idx < len(node.inodes)-1 {
					result += ","
				}
			}
		case *internalNode:
			for idx, inode := range node.inodes {
				result += strconv.Itoa(inode.key)
				if idx < len(node.inodes)-1 {
					result += ","
				}
			}
		case nil:
			result += "\n"
		default:
			panic("invalid node")
		}
		result += "/"
	}

	// clean unnecessary separator /
	newResult := ""
	splits := strings.Split(result, "\n")
	for _, each := range splits {
		newResult += strings.Trim(each, "/")
		newResult += "\n"
	}
	return strings.TrimSuffix(newResult, "\n")
}
