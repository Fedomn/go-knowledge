package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLRU(t *testing.T) {
	lruCache := Constructor(3)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Put(3, 3)
	lruCache.Put(4, 4)
	lruCache.Put(5, 5)
	lruCache.Put(6, 6)

	assert.Equal(t, -1, lruCache.Get(1))
	assert.Equal(t, -1, lruCache.Get(2))
	assert.Equal(t, -1, lruCache.Get(3))
	assert.Equal(t, 4, lruCache.Get(4))
}
