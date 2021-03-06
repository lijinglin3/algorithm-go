package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortArray(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 5}, sortArray([]int{5, 2, 3, 1}))
}
