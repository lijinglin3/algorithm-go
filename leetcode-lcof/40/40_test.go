package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLeastNumbers(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, getLeastNumbers([]int{5, 4, 3, 2, 1}, 3))
}
