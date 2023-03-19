package lc

import . "github.com/fedomn/go-knowledge/test/algorithm/lc/util"

func middleNode143(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList143(head *ListNode) *ListNode {
	var prev *ListNode
	cursor := head
	for cursor != nil {
		next := cursor.Next
		cursor.Next = prev
		prev = cursor
		cursor = next
	}
	return prev
}

func mergeList143(l, r *ListNode) *ListNode {
	dummyHeader := &ListNode{}
	cursor := dummyHeader
	// 每次 l 放在前面，r 放在后面
	for l != nil && r != nil {
		cursor.Next = l
		l = l.Next
		cursor = cursor.Next

		cursor.Next = r
		r = r.Next
		cursor = cursor.Next
	}
	if l != nil {
		cursor.Next = l
	} else if r != nil {
		cursor.Next = r
	}
	return dummyHeader.Next
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := middleNode143(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseList143(l2)
	head = mergeList143(l1, l2)
}
