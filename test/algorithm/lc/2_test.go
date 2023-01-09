package lc

import (
	. "github.com/fedomn/go-knowledge/test/algorithm/lc/util"
	"reflect"
	"testing"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tail := &ListNode{}
	head := tail
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		tail.Val = sum % 10
		carry = sum / 10
		if l1 == nil && l2 == nil {
			if carry == 0 {
				tail.Next = nil
			} else {
				tail.Next = &ListNode{Val: carry}
				tail = tail.Next
			}
		} else {
			tail.Next = &ListNode{}
			tail = tail.Next
		}
	}
	return head
}

func Test2(t *testing.T) {
	l := GenListNode([]int{2, 4, 3})
	if !reflect.DeepEqual(WalkList(l), []int{2, 4, 3}) {
		t.Fatal("GenListNode failed")
	}

	tests := []struct {
		l1     []int
		l2     []int
		expect []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	for _, test := range tests {
		l1 := GenListNode(test.l1)
		l2 := GenListNode(test.l2)
		expect := GenListNode(test.expect)
		if !reflect.DeepEqual(WalkList(addTwoNumbers(l1, l2)), WalkList(expect)) {
			t.Fatalf("expect %v, got %v", WalkList(expect), WalkList(addTwoNumbers(l1, l2)))
		}
	}
}
