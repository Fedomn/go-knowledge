package util

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenListNode(input []int) *ListNode {
	dummyHead := &ListNode{}
	cursor := dummyHead
	for i := 0; i < len(input); i++ {
		cursor.Next = &ListNode{Val: input[i]}
		cursor = cursor.Next
	}
	return dummyHead.Next
}

func WalkList(l *ListNode) []int {
	res := make([]int, 0)
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	return res
}

// 想象一颗从左向右的链表
// 翻转链表核心：
// 1. 先定义 prev 和 cursor
// 2. 处理cursor.Next
// 3. 设置 prev 和 cursor

// 总体思路，每次取出一个节点，放入新的链表中
func reverseHeadTailUtil(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	cursor := head
	// 新链表的头节点 != 要翻转的最后一个节点
	for prev != tail {
		// next作为下一次的cursor
		next := cursor.Next
		// 当前节点指向新链表的头节点
		cursor.Next = prev
		// 更新新链表的头节点
		prev = cursor
		// 更新cursor
		cursor = next
	}
	return tail, head
}

func reverseFromHeadUtil(head *ListNode) *ListNode {
	// 作为新链表头结点
	var prev *ListNode
	cursor := head
	// 新链表的头节点 != 链表的最后nil
	for cursor != nil {
		// next作为下一次的cursor
		next := cursor.Next
		// 当前节点指向新链表的头节点
		cursor.Next = prev
		// 更新新链表的头节点
		prev = cursor
		// 更新cursor
		cursor = next
	}
	return prev
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func mergeTwoList(l, r *ListNode) *ListNode {
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
