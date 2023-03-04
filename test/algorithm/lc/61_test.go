package lc

import . "github.com/fedomn/go-knowledge/test/algorithm/lc/util"

func rotateRight(head *ListNode, k int) *ListNode {
	totalLen := 0
	cursor := head
	for cursor != nil {
		cursor = cursor.Next
		totalLen++
	}
	if totalLen == 0 {
		return nil
	}

	k = k % totalLen
	if totalLen == 1 || k == 0 {
		return head
	}

	// fmt.Println(totalLen, k)

	slow, fast := head, head
	// 快慢指针之间的差值控制需要移动的节点
	cnt := 0
	for fast != nil {
		if cnt > k {
			slow = slow.Next
		}
		fast = fast.Next
		cnt++
	}
	// fmt.Println("hit, ", slow.Val)
	newHeader := slow.Next
	slow.Next = nil
	cursor = newHeader
	for cursor.Next != nil {
		cursor = cursor.Next
	}
	cursor.Next = head
	return newHeader
}
