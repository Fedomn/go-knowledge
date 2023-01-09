package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseHeadTailUtil(t *testing.T) {
	l := GenListNode([]int{1, 2, 3, 4, 5})
	// reverse 2, 3, 4
	head, tail := l.Next, l.Next.Next.Next
	assert.Equal(t, 2, head.Val)
	assert.Equal(t, 4, tail.Val)
	//fmt.Println(WalkList(head))

	head, tail = reverseHeadTailUtil(head, tail)
	assert.Equal(t, 4, head.Val)
	assert.Equal(t, 2, tail.Val)
	assert.Equal(t, []int{4, 3, 2, 5}, WalkList(head))
}

func TestReverseFromHeadUtil(t *testing.T) {
	l := GenListNode([]int{1, 2, 3, 4, 5})
	l = reverseFromHeadUtil(l)
	assert.Equal(t, []int{5, 4, 3, 2, 1}, WalkList(l))
}
