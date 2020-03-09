package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	cases := [][]int{
		{0, 1, 0, 3, 12},
		{2, 1, 0},
		{2, 1},
	}
	results := [][]int{
		{1, 3, 12, 0, 0},
		{2, 1, 0},
		{2, 1},
	}

	for i := range cases {
		MoveZeroes(cases[i])
		assert.Equal(t, cases[i], results[i])
	}
}
