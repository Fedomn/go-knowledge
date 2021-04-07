package btree

// node represents a block that contains many inodes
type node interface {
	// find finds the index of the given key, if exist return true, otherwise return false
	find(key int) (int, bool)
	insert(key int, value int, leftNode node, rightNode node) (spilledKey int, spilledValue int, spilledNode node, spilled bool)
	parent() *internalNode
	setParent(*internalNode)
	full() bool
}
