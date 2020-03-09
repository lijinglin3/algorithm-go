package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	cases := [][]int{
		{1, 2, 3},
		{1, 9, 9},
		{1, 9},
		{9},
	}
	results := [][]int{
		{1, 2, 4},
		{2, 0, 0},
		{2, 0},
		{1, 0},
	}

	for i := range cases {
		assert.Equal(t, PlusOne(cases[i]), results[i])
	}
}
