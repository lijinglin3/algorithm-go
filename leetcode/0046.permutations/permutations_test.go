package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermute(t *testing.T) {
	assert.Equal(t, [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}, permute([]int{1, 2, 3}))
}
