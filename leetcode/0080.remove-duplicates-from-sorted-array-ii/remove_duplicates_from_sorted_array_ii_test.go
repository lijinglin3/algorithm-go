package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	example := []int{1, 1, 1, 2, 2, 3}
	assert.Equal(t, 5, removeDuplicates(example))
	assert.Equal(t, []int{1, 1, 2, 2, 3}, example[:5])
}
