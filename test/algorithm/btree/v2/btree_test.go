package v2

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestInsert_Basic(t *testing.T) {
	tree := newBtree(3)
	for i := 1; i <= 11; i++ {
		tree.insert(i)
		t.Logf("After insert %d, PreOrder: %v", i, tree.preOrderTraversal())
		t.Logf("BFS: \n%s", tree.breadthFirstDraw())
	}

	expectPreOrder := []int{3, 1, 2, 6, 4, 5, 7, 8, 9, 10, 11}
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

	expectPreOrder := []int{15, 1, 9, 22, 19, 20, 30, 40, 55}
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

	expectPreOrder := []int{15, 1, 9, 10, 17, 16, 22, 18, 20, 23, 32}
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

func Benchmark_Insert_RandomInsertion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randomInsertion(b)
	}
}

func randomInsertion(b *testing.B) {
	tree := newBtree(4)
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

		tree.insert(randInt)
	}
	insertedCount := len(tree.preOrderTraversal())

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

	expectPreOrder := []int{19, 1, 15, 16, 20, 23, 28, 32}
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

	expectPreOrder := []int{15, 1, 9, 12, 13, 16, 19, 20}
	if !reflect.DeepEqual(tree.preOrderTraversal(), expectPreOrder) {
		t.Fatalf("BTree insert splitting incorrect, got preOrder: %v", tree.preOrderTraversal())
	}
}

func TestRemove_NonLeaf_GetPred(t *testing.T) {
	testData := []int{1, 15, 32, 9, 20, 16, 19, 12, 13, 38, 40}
	tree := newBtree(4)
	for _, d := range testData {
		tree.insert(d)
		t.Logf("After insert %d, PreOrder: %v", d, tree.preOrderTraversal())
		t.Logf("BFS: \n%s", tree.breadthFirstDraw())
	}

	tree.remove(16)
	t.Logf("After remove BFS: \n%s", tree.breadthFirstDraw())

	expectPreOrder := []int{15, 1, 9, 12, 13, 19, 20, 32, 38, 40}
	if !reflect.DeepEqual(tree.preOrderTraversal(), expectPreOrder) {
		t.Fatalf("BTree insert splitting incorrect, got preOrder: %v", tree.preOrderTraversal())
	}
}

func TestRemove_NonLeaf_GetSucc(t *testing.T) {
	testData := []int{1, 32, 20, 16, 19, 12, 13, 38, 40}
	tree := newBtree(4)
	for _, d := range testData {
		tree.insert(d)
		t.Logf("After insert %d, PreOrder: %v", d, tree.preOrderTraversal())
		t.Logf("BFS: \n%s", tree.breadthFirstDraw())
	}

	tree.remove(16)
	t.Logf("After remove BFS: \n%s", tree.breadthFirstDraw())

	expectPreOrder := []int{19, 1, 12, 13, 20, 32, 38, 40}
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
	t.Logf("After remove BFS: \n%s", tree.breadthFirstDraw())
	t.Log("-----------")
	tree.remove(19)
	t.Logf("After remove BFS: \n%s", tree.breadthFirstDraw())

	expectPreOrder := []int{1, 12, 13, 20, 32, 38}
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
	insertedCount := len(tree.preOrderTraversal())
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
	}

	//b.Logf("After remove BFS: \n%s", tree.breadthFirstDraw())

	if insertedCount != (count - len(duplicatedList)) {
		b.Fatalf("BTree random insert count incorrect")
	}
}

func TestCheck_Validity(t *testing.T) {
	arr := []int{450, 621, 181, 340, 945, 584, 793, 50, 156, 446, 898, 963, 79, 352, 790, 559, 995, 413, 422, 276, 986, 163, 426, 244, 742, 721, 266, 733, 303, 985, 341, 880, 500, 551, 839, 37, 30, 205, 493, 279, 314, 809, 430, 965, 868, 334, 998, 979, 69, 828, 602, 628, 94, 645, 91, 10, 285, 322, 119, 203, 336, 482, 291, 264, 59, 794, 686, 958, 14, 563, 265, 590, 715, 634, 467, 575, 608, 469, 204, 815, 732, 460, 153, 857, 508, 295, 159, 941, 417, 999, 193, 354, 675, 720, 128, 629, 849}
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
