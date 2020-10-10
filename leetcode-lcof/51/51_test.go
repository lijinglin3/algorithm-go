package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReversePairs(t *testing.T) {
	assert.Equal(t, 26, reversePairs([]int{7, 5, 6, 4, 2, 3, 1, 0}))
}
