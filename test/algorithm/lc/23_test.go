package lc

import . "github.com/fedomn/go-knowledge/test/algorithm/lc/util"

// func mergeKLists(lists []*ListNode) *ListNode {
//     var cursor *ListNode
//     for i:=0; i<len(lists); i++ {
//         cursor = mergeTwoLists(cursor, lists[i])
//     }
//     return cursor
// }

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge23(lists, 0, len(lists)-1)
}

func merge23(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	mid := (l + r) / 2
	left := merge23(lists, l, mid)
	right := merge23(lists, mid+1, r)
	return mergeTwoLists23(left, right)
}

func mergeTwoLists23(l *ListNode, r *ListNode) *ListNode {
	dummyHead := &ListNode{}
	cursor := dummyHead
	for l != nil && r != nil {
		if l.Val < r.Val {
			cursor.Next = l
			l = l.Next
		} else {
			cursor.Next = r
			r = r.Next
		}
		cursor = cursor.Next
	}
	if l != nil {
		cursor.Next = l
	} else if r != nil {
		cursor.Next = r
	}
	return dummyHead.Next
}
