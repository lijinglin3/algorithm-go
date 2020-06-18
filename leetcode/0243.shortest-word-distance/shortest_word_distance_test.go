package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestDistance(t *testing.T) {
	assert.Equal(t, 3, shortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "coding", "practice"))
	assert.Equal(t, 1, shortestDistance([]string{"a", "a", "b", "b"}, "a", "b"))
}
