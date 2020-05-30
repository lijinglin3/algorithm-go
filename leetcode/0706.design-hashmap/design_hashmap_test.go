package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	m := Constructor()

	assert.Equal(t, -1, m.Get(0))
	m.Put(0, 1)
	assert.Equal(t, 1, m.Get(0))
	m.Remove(0)
	assert.Equal(t, -1, m.Get(0))
}
