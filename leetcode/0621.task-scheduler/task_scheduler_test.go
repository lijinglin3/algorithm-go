package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeastInterval(t *testing.T) {
	assert.Equal(t, 6, leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 0))
	assert.Equal(t, 8, leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
	assert.Equal(t, 12, leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C', 'D', 'D', 'E'}, 2))
}
