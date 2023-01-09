package lc

import . "github.com/fedomn/go-knowledge/test/algorithm/lc/util"

//func detectCycle(head *ListNode) *ListNode {
//	vMap := make(map[*ListNode]bool)
//	for head != nil {
//		if _,ok:=vMap[head.Next];ok{
//			return head.Next
//		} else {
//			vMap[head] = true
//			head = head.Next
//		}
//	}
//	return nil
//}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	hasCycle := false
	for fast != nil && slow != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			hasCycle = true
			break
		}
	}
	if !hasCycle {
		return nil
	}
	// 一定有环，则fast从head开始移动1，slow从相遇点开始移动1，他们再次相遇时，为进环口
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
