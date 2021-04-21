package btree

import (
	"fmt"
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
// https://github.com/lichuang/algorithm_notes/blob/master/btree/btree.py
// https://scanftree.com/Data_Structure/deletion-in-b-tree
// https://www.programiz.com/dsa/deletion-from-a-b-tree

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

	fmt.Println(tree.breadthFirstDraw())
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

	fmt.Println(tree.breadthFirstDraw())
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

	fmt.Println(tree.breadthFirstDraw())
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

func TestDelete_LeafNode_And_BorrowFromRightSibling(t *testing.T) {
	testData := []int{1, 15, 32, 9, 20, 23, 16, 19, 28}
	tree := newTree(4)
	for _, d := range testData {
		tree.Insert(d, d)
	}

	fmt.Println(tree.breadthFirstDraw())

	tree.Delete(9)
	fmt.Println(tree.breadthFirstDraw())

	expectPreOrder := []int{16, 1, 15, 19, 20, 23, 19, 20, 28, 32}
	if !reflect.DeepEqual(tree.preOrderTraversal(), expectPreOrder) {
		t.Fatalf("BTree insert splitting incorrect, got preOrder: %v", tree.preOrderTraversal())
	}
}

func TestDelete_LeafNode_And_BorrowFromLeftSibling(t *testing.T) {
	testData := []int{1, 15, 32, 9, 20, 23, 16, 19, 28}
	tree := newTree(4)
	for _, d := range testData {
		tree.Insert(d, d)
	}

	fmt.Println(tree.breadthFirstDraw())
	tree.Delete(16)
	fmt.Println(tree.breadthFirstDraw())

	fmt.Println("----------")
	tree.Delete(19)
	fmt.Println(tree.breadthFirstDraw())

	expectPreOrder := []int{9, 1, 15, 20, 23, 15, 20, 28, 32}
	if !reflect.DeepEqual(tree.preOrderTraversal(), expectPreOrder) {
		t.Fatalf("BTree insert splitting incorrect, got preOrder: %v", tree.preOrderTraversal())
	}
}

func TestDelete_LeafNode_And_MergeRightSibling(t *testing.T) {
	testData := []int{1, 15, 32, 9, 20, 23, 16, 19, 28}
	tree := newTree(4)
	for _, d := range testData {
		tree.Insert(d, d)
	}

	fmt.Println(tree.breadthFirstDraw())
	tree.Delete(9)
	fmt.Println(tree.breadthFirstDraw())
	tree.Delete(15)
	fmt.Println(tree.breadthFirstDraw())

	fmt.Println("----------")
	tree.Delete(16)
	fmt.Println(tree.breadthFirstDraw())

	expectPreOrder := []int{23, 1, 19, 20, 28, 32}
	if !reflect.DeepEqual(tree.preOrderTraversal(), expectPreOrder) {
		t.Fatalf("BTree insert splitting incorrect, got preOrder: %v", tree.preOrderTraversal())
	}
}

func TestDelete_LeafNode_And_MergeLeftSibling(t *testing.T) {
	testData := []int{1, 15, 32, 9, 20, 23, 16, 19, 28}
	tree := newTree(4)
	for _, d := range testData {
		tree.Insert(d, d)
	}

	fmt.Println(tree.breadthFirstDraw())
	tree.Delete(9)
	fmt.Println(tree.breadthFirstDraw())
	tree.Delete(20)
	fmt.Println(tree.breadthFirstDraw())

	tree.Delete(16)
	fmt.Println(tree.breadthFirstDraw())

	fmt.Println("----------")
	tree.Delete(19)
	fmt.Println(tree.breadthFirstDraw())

	expectPreOrder := []int{28, 1, 15, 23, 32}
	if !reflect.DeepEqual(tree.preOrderTraversal(), expectPreOrder) {
		t.Fatalf("BTree insert splitting incorrect, got preOrder: %v", tree.preOrderTraversal())
	}
}

func TestDelete_InternalNode_BorrowLeft(t *testing.T) {
	testData := []int{1, 5, 9, 14, 18, 22, 28, 33, 37, 41, 46, 50, 54, 59, 61, 64, 67, 72}
	tree := newTree(4)
	for _, d := range testData {
		tree.Insert(d, d)
	}
	fmt.Println(tree.breadthFirstDraw())

	fmt.Println("----------")
	tree.Delete(9)
	fmt.Println(tree.breadthFirstDraw())

	expectPreOrderAfterDelete := []int{37, 14, 1, 5, 18, 22, 18, 28, 33, 50, 41, 46, 54, 59, 61, 54, 59, 64, 67, 72}
	if !reflect.DeepEqual(tree.preOrderTraversal(), expectPreOrderAfterDelete) {
		t.Fatalf("BTree delete incorrect, got preOrder: %v", tree.preOrderTraversal())
	}
}

func TestDelete_InternalNode_BorrowRight(t *testing.T) {
	testData := []int{1, 5, 9, 14, 18, 22, 28, 33, 37, 41, 46, 50, 54, 59, 61, 64, 67, 72}
	tree := newTree(4)
	for _, d := range testData {
		tree.Insert(d, d)
	}
	fmt.Println(tree.breadthFirstDraw())
	tree.Delete(9)
	fmt.Println(tree.breadthFirstDraw())

	fmt.Println("----------")
	tree.Delete(5)
	fmt.Println(tree.breadthFirstDraw())

	expectPreOrderAfterDelete := []int{37, 22, 1, 14, 18, 28, 33, 50, 41, 46, 54, 59, 61, 54, 59, 64, 67, 72}
	if !reflect.DeepEqual(tree.preOrderTraversal(), expectPreOrderAfterDelete) {
		t.Fatalf("BTree delete incorrect, got preOrder: %v", tree.preOrderTraversal())
	}
}

func TestDelete_InternalNode_MergeLeftRightLeafNode(t *testing.T) {
	testData := []int{1, 5, 9, 14, 18, 22, 28, 33, 37, 41, 46, 50, 54, 59, 61, 64, 67, 72}
	tree := newTree(4)
	for _, d := range testData {
		tree.Insert(d, d)
	}
	fmt.Println(tree.breadthFirstDraw())
	tree.Delete(9)
	fmt.Println(tree.breadthFirstDraw())
	tree.Delete(5)
	fmt.Println(tree.breadthFirstDraw())

	fmt.Println("----------")
	tree.Delete(14)
	fmt.Println(tree.breadthFirstDraw())

	expectPreOrderAfterDelete := []int{37, 22, 1, 18, 28, 33, 50, 41, 46, 54, 59, 61, 54, 59, 64, 67, 72}
	if !reflect.DeepEqual(tree.preOrderTraversal(), expectPreOrderAfterDelete) {
		t.Fatalf("BTree delete incorrect, got preOrder: %v", tree.preOrderTraversal())
	}
}

// FIXME
//func TestDelete_InternalNode_MergeLeftRightInternalNode(t *testing.T) {
//	testData := []int{1, 5, 9, 14, 18, 22, 28, 33, 37, 41, 46, 50, 54, 59, 61, 64, 67, 72}
//	tree := newTree(4)
//	for _, d := range testData {
//		tree.Insert(d, d)
//	}
//	fmt.Println(tree.breadthFirstDraw())
//	tree.Delete(37)
//	fmt.Println(tree.breadthFirstDraw())
//	tree.Delete(33)
//	fmt.Println(tree.breadthFirstDraw())
//
//	tree.Delete(28)
//	fmt.Println(tree.breadthFirstDraw())
//
//	tree.Delete(22)
//	fmt.Println(tree.breadthFirstDraw())
//
//	fmt.Println("----------")
//	// FIXME
//	tree.Delete(18)
//	fmt.Println(tree.breadthFirstDraw())
//
//	expectPreOrderAfterDelete := []int{37, 22, 1, 18, 28, 33, 50, 41, 46, 54, 59, 61, 54, 59, 64, 67, 72}
//	if !reflect.DeepEqual(tree.preOrderTraversal(), expectPreOrderAfterDelete) {
//		t.Fatalf("BTree delete incorrect, got preOrder: %v", tree.preOrderTraversal())
//	}
//}
