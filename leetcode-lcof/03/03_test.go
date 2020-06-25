package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRepeatNumber(t *testing.T) {
	assert.Equal(t, 2, findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
	assert.Equal(t, 0, findRepeatNumber([]int{0, 1, 2, 3, 4, 5, 6, 7}))
}
