package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumWays(t *testing.T) {
	assert.Equal(t, 1, numWays(1))
	assert.Equal(t, 21, numWays(7))
	assert.Equal(t, 720754435, numWays(92))
}
