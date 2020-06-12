package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMin(t *testing.T) {
	assert.Equal(t, 0, findMin([]int{4, 5, 6, 7, 0, 1, 2}))
}
