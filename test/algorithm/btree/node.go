package btree

// node represents a block that contains many inodes
type node interface {
	// find finds the index of the given key, if exist return true, otherwise return false
	find(key string) (int, bool)
	insert(key string, value string, leftNode node, rightNode node) (spilledKey string, spilledValue string, spilledNode node, spilled bool)
	parent() *internalNode
	setParent(*internalNode)
	full() bool
}
