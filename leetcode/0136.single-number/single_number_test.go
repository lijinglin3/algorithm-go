package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	cases := [][]int{
		{2, 2, 1},
		{4, 1, 2, 1, 2},
		{1},
		{},
	}
	results := []int{1, 4, 1, 0}

	for i := range cases {
		assert.Equal(t, results[i], singleNumber(cases[i]))
	}
}
