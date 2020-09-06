package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	cases := [][]int{
		{},
		{1},
		{7, 1, 5, 3, 6, 4},
		{1, 2, 3, 4, 5},
		{7, 6, 4, 3, 1},
	}
	results := []int{0, 0, 7, 4, 0}

	for i, s := range cases {
		assert.Equal(t, results[i], maxProfit(s))
	}
}
