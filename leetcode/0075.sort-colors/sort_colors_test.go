package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortColors(t *testing.T) {
	colors := []int{2, 0, 2, 1, 1, 0}
	sortColors(colors)
	assert.Equal(t, []int{0, 0, 1, 1, 2, 2}, colors)
}
