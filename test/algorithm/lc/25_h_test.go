package lc

import (
	. "github.com/fedomn/go-knowledge/test/algorithm/lc/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func reverseKGroup(head *ListNode, k int) *ListNode {
	newHeader := &ListNode{Next: head}
	prev, next := newHeader, &ListNode{}
	head, tail := head, newHeader

	for head != nil {
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return newHeader.Next
			}
		}
		next = tail.Next

		head, tail = reverseInternal25(head, tail)
		prev.Next = head
		tail.Next = next

		prev = tail
		head = tail.Next
	}

	return newHeader.Next
}

func reverseInternal25(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	cursor := head
	for prev != tail {
		next := cursor.Next
		cursor.Next = prev
		prev = cursor
		cursor = next
	}
	return tail, head
}

func Test25(t *testing.T) {
	l := GenListNode([]int{1, 2, 3, 4, 5})
	l = reverseKGroup(l, 2)
	assert.Equal(t, []int{2, 1, 4, 3, 5}, WalkList(l))
}
