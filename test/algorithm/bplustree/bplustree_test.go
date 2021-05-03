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

func TestRemove_Leaf_BorrowFromNext(t *testing.T) {
	testData := []int{1, 15, 32, 9, 20, 23, 16, 19, 28}
	tree := newBtree(4)
	for _, d := range testData {
		tree.insert(d)
		t.Logf("After insert %d, PreOrder: %v", d, tree.preOrderTraversal())
		t.Logf("BFS: \n%s", tree.breadthFirstDraw())
	}

	tree.remove(9)
	t.Logf("After remove BFS: \n%s", tree.breadthFirstDraw())

	expectPreOrder := []int{19, 1, 15, 16, 19, 20, 23, 28, 32}
	if !reflect.DeepEqual(tree.preOrderTraversal(), expectPreOrder) {
		t.Fatalf("BTree insert splitting incorrect, got preOrder: %v", tree.preOrderTraversal())
	}
}

func TestRemove_Leaf_BorrowFromPrev(t *testing.T) {
	testData := []int{1, 15, 32, 9, 20, 16, 19, 12, 13}
	tree := newBtree(4)
	for _, d := range testData {
		tree.insert(d)
		t.Logf("After insert %d, PreOrder: %v", d, tree.preOrderTraversal())
		t.Logf("BFS: \n%s", tree.breadthFirstDraw())
	}

	tree.remove(32)
	t.Logf("After remove BFS: \n%s", tree.breadthFirstDraw())
	tree.remove(20)
	t.Logf("After remove BFS: \n%s", tree.breadthFirstDraw())

	expectPreOrder := []int{15, 1, 9, 12, 13, 15, 16, 19}
	if !reflect.DeepEqual(tree.preOrderTraversal(), expectPreOrder) {
		t.Fatalf("BTree insert splitting incorrect, got preOrder: %v", tree.preOrderTraversal())
	}
}

func TestRemove_NonLeaf_Merge(t *testing.T) {
	testData := []int{1, 32, 20, 16, 19, 12, 13, 38}
	tree := newBtree(4)
	for _, d := range testData {
		tree.insert(d)
		t.Logf("After insert %d, PreOrder: %v", d, tree.preOrderTraversal())
		t.Logf("BFS: \n%s", tree.breadthFirstDraw())
	}

	tree.remove(16)
	t.Logf("After remove %d, BFS: \n%s", 16, tree.breadthFirstDraw())
	tree.remove(19)
	t.Logf("After remove %d, BFS: \n%s", 19, tree.breadthFirstDraw())
	t.Log("-----------")
	tree.remove(1)
	t.Logf("After remove %d, BFS: \n%s", 1, tree.breadthFirstDraw())

	expectPreOrder := []int{12, 13, 20, 32, 38}
	if !reflect.DeepEqual(tree.preOrderTraversal(), expectPreOrder) {
		t.Fatalf("BTree insert splitting incorrect, got preOrder: %v", tree.preOrderTraversal())
	}
}

func Benchmark_Remove_RandomRemoving(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randomRemoving(b)
	}
}

func randomRemoving(b *testing.B) {
	tree := newBtree(4)
	rand.Seed(time.Now().UnixNano())
	existMap := make(map[int]bool, 0)
	duplicatedList := make([]int, 0)
	insertedList := make([]int, 0)
	count := 100
	for i := 0; i < count; i++ {
		randInt := rand.Intn(1000)
		if existMap[randInt] {
			duplicatedList = append(duplicatedList, randInt)
		} else {
			existMap[randInt] = true
			insertedList = append(insertedList, randInt)
		}

		tree.insert(randInt)
	}
	firstKey := tree.firstKey()
	scanKeys := tree.scanKeysFrom(firstKey)
	insertedCount := len(scanKeys)
	b.Logf("%v", insertedList)

	if insertedCount != (count - len(duplicatedList)) {
		b.Logf("BTree random insert count incorrect\n")
		b.Logf("insertedCount: %d, duplicatedList: %d, sum: %d, existMap: %d\n", insertedCount, len(duplicatedList), insertedCount+len(duplicatedList), len(existMap))
		b.Logf("duplicatedValues: %v\n", filterDuplicatedValues(tree.preOrderTraversal()))
		b.Logf("BFS:\n%s", tree.breadthFirstDraw())
		b.Fatal("Done")
	}
	b.Logf("insert done, BFS:\n%s", tree.breadthFirstDraw())

	b.Log("start random removing")
	for _, key := range insertedList {
		tree.remove(key)
		if tree.root == nil {
			continue
		}
		tree.checkValidity()
		tree.checkNextPointer()
		b.Logf("After remove %d BFS: \n%s", key, tree.breadthFirstDraw())
	}

	b.Logf("After remove done BFS: \n%s", tree.breadthFirstDraw())

	if len(tree.preOrderTraversal()) != 0 {
		b.Fatalf("After BTree random remove, tree should be empty")
	}
}

func TestCheck_Validity(t *testing.T) {
	arr := []int{366, 901, 192, 92, 805, 997, 767, 30, 73, 965, 940, 896, 679, 62, 735, 81, 276, 455, 479, 880, 460, 148, 507, 144, 770, 955, 492, 645, 130, 506, 435, 704, 570, 639, 547, 332, 629, 993, 641, 623, 643, 403, 614, 783, 786, 808, 609, 732, 936, 368, 70, 664, 370, 253, 810, 429, 272, 88, 396, 524, 362, 401, 304, 230, 232, 804, 402, 252, 471, 200, 995, 162, 223, 991, 456, 493, 625, 329, 964, 225, 642, 543, 267, 859, 453, 605, 448, 583, 968, 875, 31}
	tree := newBtree(4)
	for i := 0; i < len(arr); i++ {
		tree.insert(arr[i])
	}
	t.Logf("initialized:\n%s", tree.breadthFirstDraw())

	for i := 0; i < len(arr); i++ {
		tree.remove(arr[i])
		t.Logf("after remove %d, BFS:\n%s", arr[i], tree.breadthFirstDraw())
		tree.checkValidity()
	}
}
