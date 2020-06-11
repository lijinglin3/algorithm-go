package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetZeroes(t *testing.T) {
	cases := [][][]int{
		{},
		{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		},
		{
			{0, 1, 2, 0},
			{3, 4, 5, 2},
			{1, 3, 1, 5},
		},
	}
	results := [][][]int{
		{},
		{
			{1, 0, 1},
			{0, 0, 0},
			{1, 0, 1},
		},
		{
			{0, 0, 0, 0},
			{0, 4, 5, 0},
			{0, 3, 1, 0},
		},
	}

	for i := range cases {
		setZeroes(cases[i])
		assert.Equal(t, results[i], cases[i])
	}
}
