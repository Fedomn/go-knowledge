package btree

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
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
		tree.Insert(i, i)
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
		tree.Insert(d, d)
		t.Logf("After insert %d, PreOrder: %v", d, tree.preOrderTraversal())
	}

	expectPreOrder := []int{15, 1, 9, 20, 22}
	if !reflect.DeepEqual(tree.preOrderTraversal(), expectPreOrder) {
		t.Fatalf("BTree insert splitting incorrect, got preOrder: %v", tree.preOrderTraversal())
	}
}

func TestInsert_IntermediateInternalInode(t *testing.T) {
	testData := []int{1, 15, 32, 9, 20, 22, 23, 16, 17}
	tree := newTree(4)
	for _, d := range testData {
		tree.Insert(d, d)
		t.Logf("After insert %d, PreOrder: %v", d, tree.preOrderTraversal())
	}

	expectPreOrder := []int{15, 1, 9, 16, 17, 20, 16, 17, 22, 23, 22, 32}
	if !reflect.DeepEqual(tree.preOrderTraversal(), expectPreOrder) {
		t.Fatalf("BTree insert splitting incorrect, got preOrder: %v", tree.preOrderTraversal())
	}
}

func TestInsert_SameKey(t *testing.T) {
	tree := newTree(10)
	tree.Insert(1, 1)
	tree.Insert(1, 10)
	searchVal, exist := tree.Search(1)
	if !(exist && searchVal == 10) {
		t.Fatalf("BTree insert same key but value not replace, got: %v, except: %v", searchVal, 10)
	}
}

func TestInsert_RandomInsertion(t *testing.T) {
	tree := newTree(10)
	rand.Seed(time.Now().UnixNano())
	existMap := make(map[int]bool, 0)
	duplicatedList := make([]int, 0)
	for i := 0; i < 300; i++ {
		randInt := rand.Intn(1000)
		if existMap[randInt] {
			duplicatedList = append(duplicatedList, randInt)
		} else {
			existMap[randInt] = true
		}

		tree.Insert(randInt, randInt)
	}
	insertedCount := len(unique(tree.preOrderTraversal()))
	t.Logf("inserted count: %d", insertedCount)
	t.Logf("duplicated keys: %v", duplicatedList)

	if insertedCount != (300 - len(duplicatedList)) {
		t.Fatalf("BTree random insert count incorrect")
	}
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	var list []int
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
