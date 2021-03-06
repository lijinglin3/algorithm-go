package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	example := []int{4, 6, 8, 0, 0, 0, 0, 0}
	merge(example, 3, []int{1, 2, 3, 5, 7}, 5)
	assert.Equal(t, example, []int{1, 2, 3, 4, 5, 6, 7, 8})
}
