package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMoves(t *testing.T) {
	assert.Equal(t, 3, minMoves([]int{1, 2, 3}))
}
