package lc

import . "github.com/fedomn/go-knowledge/test/algorithm/lc/util"
import (
	"reflect"
	"testing"
)

//func reverseList(head *ListNode) *ListNode {
//	var prev *ListNode
//	cursor := head
//	for cursor != nil {
//		next := cursor.Next
//		cursor.Next = prev
//		prev = cursor
//		cursor = next
//	}
//	return prev
//}

func reverseList(head *ListNode) *ListNode {
	return reverseListInternal(head, nil)
}

func reverseListInternal(head *ListNode, next *ListNode) *ListNode {
	if head == nil {
		return next
	}

	oriNext := head.Next
	head.Next = next
	next = head
	head = oriNext
	return reverseListInternal(head, next)
}

func Test206(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1}, []int{1}},
	}

	for _, test := range tests {
		got := WalkList(reverseList(GenListNode(test.input)))
		if !reflect.DeepEqual(got, test.expect) {
			t.Fatalf("input: %v, expect: %v, got: %v", test.input, test.expect, got)
		}
	}
}
