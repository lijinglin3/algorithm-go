package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	assert.Equal(t, 0, search([]int{}, 10))
	assert.Equal(t, 1, search([]int{10}, 10))
	assert.Equal(t, 0, search([]int{11}, 10))

	assert.Equal(t, 1, search([]int{5, 7, 7, 8, 8, 10}, 10))
	assert.Equal(t, 2, search([]int{5, 7, 7, 8, 8, 10}, 8))
	assert.Equal(t, 0, search([]int{5, 7, 7, 8, 8, 10}, 6))

	assert.Equal(t, 1, search([]int{5, 7, 7, 8, 8, 10, 11}, 10))
	assert.Equal(t, 2, search([]int{5, 7, 7, 8, 8, 10, 11}, 8))
	assert.Equal(t, 0, search([]int{5, 7, 7, 8, 8, 10, 11}, 6))
}
