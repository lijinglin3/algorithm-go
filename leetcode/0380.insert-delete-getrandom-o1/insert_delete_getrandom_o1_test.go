package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	set := Constructor()
	assert.Equal(t, false, set.Remove(0))
	assert.Equal(t, true, set.Insert(0))
	assert.Equal(t, false, set.Insert(0))
	assert.Equal(t, 0, set.GetRandom())
	assert.Equal(t, true, set.Remove(0))
	assert.Equal(t, true, set.Insert(0))
}
