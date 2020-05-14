package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanThreePartsEqualSum(t *testing.T) {
	cases := [][]int{
		{1, 2, 3, 4, 5, 6, 7},
		{3, 3, 6, 5, -2, 2, 5, 1, -9, 4},
		{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1},
		{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1},
		{18, 12, -18, 18, -19, -1, 10, 10},
		{10, -10, -10, 10},
		{10, -10, 10, -10, -10, 10},
	}
	results := []bool{false, true, false, true, true, false, true}

	for i := range cases {
		assert.Equal(t, results[i], canThreePartsEqualSum(cases[i]))
	}
}
