package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThirdMax(t *testing.T) {
	assert.Equal(t, 1, thirdMax([]int{0, 1}))
	assert.Equal(t, 1, thirdMax([]int{3, 2, 1}))
	assert.Equal(t, 3, thirdMax([]int{3, 1, 1}))
	assert.Equal(t, 3, thirdMax([]int{0, 5, 5, 4, 4, 3, 3, 2, 2, 1, 1}))
}
