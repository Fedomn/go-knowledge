package lc

import . "github.com/fedomn/go-knowledge/test/algorithm/lc/util"

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	slow, fast := dummyHead, head
	for fast != nil {
		if fast.Val == val {
			slow.Next = fast.Next
			fast = fast.Next
		} else {
			slow = slow.Next
			fast = fast.Next
		}
	}
	return dummyHead.Next
}
