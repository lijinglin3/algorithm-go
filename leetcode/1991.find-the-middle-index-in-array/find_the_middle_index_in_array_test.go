package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMiddleIndex(t *testing.T) {
	assert.Equal(t, -1, findMiddleIndex([]int{2, 5}))
	assert.Equal(t, 3, findMiddleIndex([]int{2, 3, -1, 8, 4}))
}
