package btree

import (
	"reflect"
	"testing"
)

// https://en.wikipedia.org/wiki/B%2B_tree
// A B+ tree consists of a root, internal nodes and leaves
// block-oriented storage data-structure
// unlike binary search trees, B+ trees have very high fanout, which reduces the number of I/O operations required to find an element in the tree.

// https://www.cs.usfca.edu/~galles/visualization/BPlusTree.html
// http://www.cburch.com/cs/340/reading/btree/index.html
// https://github.com/xiang90/bplustree
// https://github.com/abhishekchaturvedi/bplustree
// https://github.com/collinglass/bptree

// Important:
// verify tools: https://www.cs.usfca.edu/~galles/visualization/BTree.html
func TestInsert_Basic(t *testing.T) {
	tree := newTree(3)
	for i := 1; i <= 11; i++ {
		err := tree.Insert(i, i)
		if err != nil {
			t.Fatalf("BTree insert error %v", err)
		}
		t.Logf("After insert %d, PreOrder: %v", i, tree.preOrderTraversal())
	}

	expectPreOrder := []int{4, 2, 1, 3, 6, 5, 7, 8, 6, 5, 7, 10, 9, 11}
	if !reflect.DeepEqual(tree.preOrderTraversal(), expectPreOrder) {
		t.Fatalf("BTree insert splitting incorrect, got preOrder: %v", tree.preOrderTraversal())
	}
}

func TestInsert_IntermediateLeafInode(t *testing.T) {
	testData := []int{1, 15, 22, 9, 20}
	tree := newTree(4)
	for _, d := range testData {
		err := tree.Insert(d, d)
		if err != nil {
			t.Fatalf("BTree insert error %v", err)
		}
		t.Logf("After insert %d, PreOrder: %v", d, tree.preOrderTraversal())
	}

	expectPreOrder := []int{15, 1, 9, 20, 22}
	if !reflect.DeepEqual(tree.preOrderTraversal(), expectPreOrder) {
		t.Fatalf("BTree insert splitting incorrect, got preOrder: %v", tree.preOrderTraversal())
	}
}
