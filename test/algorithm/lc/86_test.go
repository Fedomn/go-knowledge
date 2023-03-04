package lc

import (
	"fmt"
	. "github.com/fedomn/go-knowledge/test/algorithm/lc/util"
)

// 把小的元素单独放到一个链表中，大的元素保留在原始链表，原始链表需要删除元素
func partition86(head *ListNode, x int) *ListNode {
	dummyHead := &ListNode{Val: -1 << 31, Next: head}
	prev, cursor := dummyHead, head
	var largeCursor *ListNode
	largeHead := &ListNode{Val: -1 << 31}
	for cursor != nil {
		// walk(dummyHead.Next)
		// walk(largeHead.Next)
		// fmt.Println("hit, ", cursor.Val)
		if cursor.Val >= x {
			next := cursor.Next
			prev.Next = next
			cursor.Next = nil
			if largeCursor == nil {
				largeCursor = cursor
				largeHead.Next = largeCursor
			} else {
				largeCursor.Next = cursor
				largeCursor = largeCursor.Next
			}
			cursor = next
		} else {
			prev = prev.Next
			cursor = cursor.Next
		}
	}
	// merge
	prev.Next = largeHead.Next
	return dummyHead.Next
}

func walk(h *ListNode) {
	res := make([]int, 0)
	for h != nil {
		res = append(res, h.Val)
		h = h.Next
	}
	fmt.Println(res)
}
