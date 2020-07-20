package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	assert.Equal(t, -1, search([]int{0}, 2))
	assert.Equal(t, 0, search([]int{0}, 0))
	assert.Equal(t, -1, search([]int{-1, 0, 3, 5, 9, 12}, 2))
	assert.Equal(t, 4, search([]int{-1, 0, 3, 5, 9, 12}, 9))
}
