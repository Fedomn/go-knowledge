package lc

import . "github.com/fedomn/go-knowledge/test/algorithm/lc/util"

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	cursor := dummyHead
	for cursor.Next != nil && cursor.Next.Next != nil {
		first := cursor
		second := cursor.Next
		third := cursor.Next.Next
		first.Next = third
		second.Next = third.Next
		third.Next = second
		// fmt.Print(walk(cursor))
		cursor = second
		// fmt.Println(",", walk(cursor))
	}
	return dummyHead.Next
}
