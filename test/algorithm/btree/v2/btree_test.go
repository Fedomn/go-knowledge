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

	expectPreOrder := []int{3, 1, 2, 4, 5, 6, 4, 5, 7, 8, 9, 10, 11}
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

	expectPreOrder := []int{15, 1, 9, 19, 20, 22, 19, 20, 30, 40, 55}
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

	expectPreOrder := []int{15, 1, 9, 10, 16, 17, 16, 18, 20, 22, 18, 20, 23, 32}
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

func TestInsert_RandomInsertion(t *testing.T) {
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
	insertedCount := len(unique(tree.preOrderTraversal()))
	t.Logf("inserted count: %d", insertedCount)
	t.Logf("duplicated keys: %v", duplicatedList)

	t.Logf("BFS:\n%s", tree.breadthFirstDraw())

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
