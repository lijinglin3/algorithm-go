package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	cases := [][]int{
		{1, 2, 3, 1},
		{1, 2, 3, 4},
		{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
		{1},
		{},
	}
	results := []bool{
		true,
		false,
		true,
		false,
		false,
	}

	for i := range cases {
		assert.Equal(t, results[i], ContainsDuplicate(cases[i]))
	}
}
