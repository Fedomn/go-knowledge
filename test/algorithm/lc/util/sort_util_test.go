package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{5, 4, 2, 3, 1}
	BubbleSort(nums)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, nums)
}

func TestQuickSort(t *testing.T) {
	nums := []int{5, 4, 2, 3, 1}
	QuickSort(nums)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, nums)
}
