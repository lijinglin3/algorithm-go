package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	cases := [][2][]int{
		{{1, 2, 2, 1}, {2, 2}},
		{{4, 9, 5}, {9, 4, 9, 8, 4}},
		{{0}, {1}},
		{{}, {}},
	}
	results := [][]int{
		{2, 2},
		{4, 9},
		{},
		{},
	}

	for i := range cases {
		assert.ElementsMatch(t, results[i], Intersection(cases[i][0], cases[i][1]))
	}
}
