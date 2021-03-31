package main

import (
	"fmt"
	"sort"
	"testing"
)

func Test_SortSearch(t *testing.T) {
	nodes := []int{1, 5, 2, 4, 3, 1, 2, 22222}

	find := 2
	sort.Ints(nodes)
	idx, present := searchFirstLargerIdx(nodes, find)
	fmt.Printf("nodes: %v, find: %d, idx: %d, present: %t\n", nodes, find, idx, present)

	find = 12345
	sort.Ints(nodes)
	idx, present = searchFirstLargerIdx(nodes, find)
	fmt.Printf("nodes: %v, find: %d, idx: %d, present: %t\n", nodes, find, idx, present)

	find = -2
	sort.Ints(nodes)
	idx, present = searchFirstLargerIdx(nodes, find)
	fmt.Printf("nodes: %v, find: %d, idx: %d, present: %t\n", nodes, find, idx, present)

	find = 2
	nodes = []int{}
	idx, present = searchFirstLargerIdx(nodes, find)
	fmt.Printf("nodes: %v, find: %d, idx: %d, present: %t\n", nodes, find, idx, present)
}

func searchFirstLargerIdx(nodes []int, find int) (int, bool) {
	idx := sort.Search(len(nodes), func(i int) bool {
		return nodes[i] >= find
	})
	return idx, idx < len(nodes) && nodes[idx] == find
}
