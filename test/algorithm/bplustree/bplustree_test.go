package bplustree

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestInsert_Basic(t *testing.T) {
	tree := newBtree(2)
	for i := 1; i <= 11; i++ {
		tree.insert(i)
		t.Logf("After insert %d, PreOrder: %v", i, tree.preOrderTraversal())
		t.Logf("BFS: \n%s", tree.breadthFirstDraw())
		tree.checkValidity()
	}

	expectPreOrder := []int{3, 2, 1, 2, 5, 4, 3, 4, 7, 6, 5, 6, 8, 7, 9, 8, 9, 10, 11}
	if !reflect.DeepEqual(tree.preOrderTraversal(), expectPreOrder) {
		t.Fatalf("BTree insert splitting incorrect, got preOrder: %v", tree.preOrderTraversal())
	}
}

func TestInsert_IntermediateLeafInode(t *testing.T) {
	testData := []int{1, 15, 22, 9, 20, 30, 40, 55, 19}
	tree := newBtree(2)
	for _, d := range testData {
		tree.insert(d)
		t.Logf("After insert %d, PreOrder: %v", d, tree.preOrderTraversal())
		t.Logf("BFS: \n%s", tree.breadthFirstDraw())
	}

	expectPreOrder := []int{20, 15, 1, 9, 15, 19, 22, 20, 30, 22, 30, 40, 55}
	if !reflect.DeepEqual(tree.preOrderTraversal(), expectPreOrder) {
		t.Fatalf("BTree insert splitting incorrect, got preOrder: %v", tree.preOrderTraversal())
	}
}

func TestInsert_IntermediateInternalInode(t *testing.T) {
	testData := []int{1, 15, 32, 9, 20, 22, 23, 16, 17, 10, 18}
	tree := newBtree(2)
	for _, d := range testData {
		tree.insert(d)
		t.Logf("After insert %d, PreOrder: %v", d, tree.preOrderTraversal())
		t.Logf("BFS: \n%s", tree.breadthFirstDraw())
	}

	expectPreOrder := []int{20, 15, 1, 9, 10, 16, 15, 16, 17, 18, 22, 20, 22, 23, 32}
	if !reflect.DeepEqual(tree.preOrderTraversal(), expectPreOrder) {
		t.Fatalf("BTree insert splitting incorrect, got preOrder: %v", tree.preOrderTraversal())
	}
}

func TestInsert_SameKey(t *testing.T) {
	tree := newBtree(2)
	tree.insert(1)
	tree.insert(1)
	t.Logf("BFS:\n%s", tree.breadthFirstDraw())

	expectPreOrder := []int{1}
	if !reflect.DeepEqual(tree.preOrderTraversal(), expectPreOrder) {
		t.Fatalf("BTree insert splitting incorrect, got preOrder: %v", tree.preOrderTraversal())
	}
}

func TestInsert_ScanKeys(t *testing.T) {
	tree := newBtree(2)
	for i := 1; i <= 11; i++ {
		tree.insert(i)
		t.Logf("After insert %d, PreOrder: %v", i, tree.preOrderTraversal())
		t.Logf("BFS: \n%s", tree.breadthFirstDraw())
		tree.checkValidity()
	}
	scanKeys := tree.scanKeysFrom(1)
	expectPreOrder := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	if !reflect.DeepEqual(scanKeys, expectPreOrder) {
		t.Fatalf("BTree insert splitting incorrect, got preOrder: %v", tree.preOrderTraversal())
	}
}

func Benchmark_Insert_RandomInsertion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randomInsertion(b)
	}
}

func randomInsertion(b *testing.B) {
	tree := newBtree(4)
	rand.Seed(time.Now().UnixNano())
	existMap := make(map[int]bool, 0)
	insertedList := make([]int, 0)
	duplicatedList := make([]int, 0)
	for i := 0; i < 300; i++ {
		randInt := rand.Intn(1000)
		if existMap[randInt] {
			duplicatedList = append(duplicatedList, randInt)
		} else {
			existMap[randInt] = true
			insertedList = append(insertedList, randInt)
		}
		tree.insert(randInt)
		// b.Logf("After insert %d, BFS: \n%s", randInt, tree.breadthFirstDraw())
		tree.checkValidity()
		tree.checkNextPointer()
	}

	sort.Ints(insertedList)
	scanKeys := tree.scanKeysFrom(insertedList[0])
	insertedCount := len(scanKeys)

	if insertedCount != (300 - len(duplicatedList)) {
		b.Logf("BTree random insert count incorrect\n")
		b.Logf("insertedCount: %d, duplicatedList: %d, sum: %d, existMap: %d\n", insertedCount, len(duplicatedList), insertedCount+len(duplicatedList), len(existMap))
		b.Logf("duplicatedValues: %v\n", filterDuplicatedValues(tree.preOrderTraversal()))
		b.Logf("BFS:\n%s", tree.breadthFirstDraw())
		b.Fatal("Done")
	}
}

func filterDuplicatedValues(intSlice []int) []int {
	keys := make(map[int]bool)
	var duplicatedValues []int
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
		} else {
			duplicatedValues = append(duplicatedValues, entry)
		}
	}
	return duplicatedValues
}
