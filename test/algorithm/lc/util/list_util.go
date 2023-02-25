package util

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenListNode(input []int) *ListNode {
	cursor := &ListNode{}
	root := cursor
	for i := 0; i < len(input); i++ {
		cursor.Val = input[i]
		if i+1 > len(input)-1 {
			return root
		} else {
			cursor.Next = &ListNode{}
			cursor = cursor.Next
		}
	}
	return root
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
