package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	cases := [][]int{
		{1, 2, 3, 4, 5, 6, 7},
		{-1, -100, 3, 99},
		{1},
	}
	steps := []int{3, 2, 10}
	results := [][]int{
		{5, 6, 7, 1, 2, 3, 4},
		{3, 99, -1, -100},
		{1},
	}

	for i := range cases {
		rotate(cases[i], steps[i])
		assert.Equal(t, results[i], cases[i])
	}
}
