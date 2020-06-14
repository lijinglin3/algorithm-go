package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	example := []byte("hello")
	reverseString(example)
	assert.Equal(t, []byte("olleh"), example)
}
