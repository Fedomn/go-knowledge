package lc

import (
	. "github.com/fedomn/go-knowledge/test/algorithm/lc/util"
	"reflect"
	"testing"
)

// 快慢指针，倒数第n个：slow比fast慢n-1个
// prevSlow作为慢指针之前的一个节点, slow作为最后要删除的节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Next: head}
	prevSlow, fast, slow := dummyNode, head, head
	acc := 0
	for fast != nil {
		if acc == n {
			prevSlow = slow
			slow = slow.Next
		} else {
			acc++
		}
		fast = fast.Next
	}
	// 开始删除slow节点
	prevSlow.Next = slow.Next
	return dummyNode.Next
}

func Test19(t *testing.T) {
	tests := []struct {
		input  []int
		n      int
		expect []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1}, 1, []int{}},
		{[]int{1, 2}, 1, []int{1}},
	}

	for _, test := range tests {
		got := WalkList(removeNthFromEnd(GenListNode(test.input), test.n))
		if !reflect.DeepEqual(got, test.expect) {
			t.Fatalf("input: %v, expect: %v, got: %v", test.input, test.expect, got)
		}
	}
}
