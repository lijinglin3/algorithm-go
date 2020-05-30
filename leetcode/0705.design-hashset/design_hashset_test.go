package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	hs := Constructor()

	hs.Add(0)
	assert.Equal(t, true, hs.Contains(0))
	hs.Remove(0)
	assert.Equal(t, false, hs.Contains(0))
}
