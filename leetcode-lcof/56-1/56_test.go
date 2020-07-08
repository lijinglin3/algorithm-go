package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumbers(t *testing.T) {
	assert.ElementsMatch(t, []int{2, 10}, singleNumbers([]int{1, 2, 10, 4, 1, 4, 3, 3}))
}
