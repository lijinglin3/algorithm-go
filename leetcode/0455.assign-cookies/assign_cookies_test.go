package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindContentChildren(t *testing.T) {
	assert.Equal(t, 1, findContentChildren([]int{1, 2, 3}, []int{1, 1}))
}
