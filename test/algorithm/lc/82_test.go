package lc

import . "github.com/fedomn/go-knowledge/test/algorithm/lc/util"

func deleteDuplicates82(head *ListNode) *ListNode {
	dummyHead := &ListNode{Val: -1 << 31, Next: head}
	cursor := dummyHead
	for cursor.Next != nil && cursor.Next.Next != nil {
		// fmt.Println("cursor, ", cursor.Val)
		if cursor.Next.Val == cursor.Next.Next.Val {
			// 找到最后一个不相同的节点
			start := cursor.Next.Next
			for start.Next != nil {
				if start.Val == start.Next.Val {
					start = start.Next
				} else {
					break
				}
			}
			// fmt.Println("hit, end, ", start.Val)
			start = start.Next
			cursor.Next = start
		} else {
			cursor = cursor.Next
		}
	}
	return dummyHead.Next
}
