package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLFU(t *testing.T) {
	lfu := ConstructorLFU(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	assert.Equal(t, 1, lfu.Get(1))
	lfu.Put(3, 3)
	assert.Equal(t, -1, lfu.Get(2))
	assert.Equal(t, 3, lfu.Get(3))
}
