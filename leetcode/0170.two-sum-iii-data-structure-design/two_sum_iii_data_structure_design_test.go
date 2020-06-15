package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	s := Constructor()

	s.Add(0)
	s.Add(-1)
	s.Add(-1)
	s.Add(0)
	s.Add(3)
	s.Add(10)

	assert.Equal(t, true, s.Find(-2))
	assert.Equal(t, true, s.Find(0))
	assert.Equal(t, true, s.Find(-1))
	assert.Equal(t, false, s.Find(1))
	assert.Equal(t, false, s.Find(5))
	assert.Equal(t, false, s.Find(-9))
}
