package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDuplicate(t *testing.T) {
	assert.Equal(t, 2, findDuplicate([]int{1, 3, 4, 2, 2}))
	assert.Equal(t, 3, findDuplicate([]int{3, 1, 3, 4, 2}))
}
