package btree

type BPlusTree struct {
	root   node
	degree int
}

func newTree(degree int) *BPlusTree {
	return &BPlusTree{
		degree: degree,
	}
}

func (tree *BPlusTree) Insert(key string, value string) {
	if tree.root == nil {
		leafNode := newLeafNode(tree.degree)
		leafNode.insert(key, value)
		tree.root = leafNode
		return
	}

	internalNodeIdx, leafNode, leafInode := tree.search(tree.root, key)
	spilledKey, spilledNode, spilled := leafNode.insert(key, value)
	if !spilled {
		return
	}

	parent := leafNode.parent()
	//var midNode node
	//midNode = leafNode
	for false {

		pSpilledKey, pSpilledNode, pSpilled := parent.insert(spilledKey, spilledNode)
		if !pSpilled {
			return
		}

		pParent := parent.parent()
		if pParent == nil {
			// root node
		}

		pParent.insert()

	}

}

func (tree BPlusTree) Search(key string) (string, bool) {
	_, _, leafInode := tree.search(tree.root, key)
	if leafInode == nil {
		return "", false
	}
	return leafInode.value, true
}

func (tree *BPlusTree) search(n node, key string) (inodeIdx int, foundNode node) {
	cursor := n
	inodeIdx = -1
	for {
		switch node := cursor.(type) {
		case *leafNode:
			idx, ok := node.find(key)
			if !ok {
				return internalNodeIdx, node, nil
			}
			return internalNodeIdx, node, &node.inodes[idx]
		case *internalNode:
			// found inodeIdx associated key always >= given key
			idx, ok := node.find(key)
			if ok {
				return idx, node
			}

			node.inodes[idx]

			cursor = node.inodes[idx].right
			inodeIdx = idx
		default:
			panic("invalid node")
		}
	}
}
