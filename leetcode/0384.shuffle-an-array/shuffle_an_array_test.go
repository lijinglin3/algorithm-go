package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	c := Constructor([]int{1, 2, 3})
	assert.ElementsMatch(t, []int{1, 2, 3}, c.Shuffle())
	assert.Equal(t, []int{1, 2, 3}, c.Reset())
}
