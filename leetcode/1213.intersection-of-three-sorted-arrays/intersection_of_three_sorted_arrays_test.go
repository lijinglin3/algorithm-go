package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArraysIntersection(t *testing.T) {
	assert.Equal(t, []int{1, 5},
		arraysIntersection(
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 5, 7, 9},
			[]int{1, 3, 4, 5, 8},
		),
	)
}
