package btree

// node represents a block that contains many inodes
type node interface {
	// find finds the index of the given key, if exist return true, otherwise return false
	find(key string) (int, bool)
	//parent() *internalNode
	//setParent(*internalNode)
	full() bool
}
