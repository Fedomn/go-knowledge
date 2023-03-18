package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeapMax(t *testing.T) {
	nums := []int{3, 8, 1, 2, 7}
	heap := NewMaxHeap()
	for _, num := range nums {
		heap.insert(num)
	}
	assert.Equal(t, heap.removeMax(), 8)
	assert.Equal(t, heap.removeMax(), 7)

}

func TestHeapMin(t *testing.T) {
	nums := []int{3, 8, 1, 2, 7}
	heap := NewMinHeap()
	for _, num := range nums {
		heap.insert(num)
	}
	//[0, 1, 2, 7, 3, 8]
	assert.Equal(t, heap.removeMin(), 1)
	assert.Equal(t, heap.removeMin(), 2)
}
