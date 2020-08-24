package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	lru := Constructor(3)

	assert.Equal(t, -1, lru.Get(1))

	lru.Put(1, 1)
	assert.Equal(t, 1, lru.Get(1))
	assert.Equal(t, -1, lru.Get(2))

	lru.Put(2, 2)
	lru.Put(3, 3)
	lru.Put(4, 4)
	assert.Equal(t, -1, lru.Get(1))
	assert.Equal(t, 2, lru.Get(2))
	assert.Equal(t, 3, lru.Get(3))
	assert.Equal(t, 4, lru.Get(4))

	lru.Put(3, 30)
	assert.Equal(t, 30, lru.Get(3))

	lru.Put(5, 5)
	lru.Put(6, 6)
	assert.Equal(t, 5, lru.Get(5))
	assert.Equal(t, 6, lru.Get(6))
	assert.Equal(t, 30, lru.Get(3))
}
