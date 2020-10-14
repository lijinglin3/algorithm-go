package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumWays(t *testing.T) {
	assert.Equal(t, 0, numWays(0, 2))
	assert.Equal(t, 2, numWays(1, 2))
	assert.Equal(t, 4, numWays(2, 2))
	assert.Equal(t, 6, numWays(3, 2))
}
