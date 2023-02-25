package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 8, 10}
	assert.Equal(t, 3, BinarySearchUtil(nums, 8))
	assert.Equal(t, -1, BinarySearchUtil(nums, 6))
}

func TestBinaryRightMost(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 8, 10}
	assert.Equal(t, 5, BinaryRightMostUtil(nums, 8))
	assert.Equal(t, 3, BinaryLeftMostUtil(nums, 8))
}
