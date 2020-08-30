package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncreasingTriplet(t *testing.T) {
	assert.Equal(t, true, increasingTriplet([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, false, increasingTriplet([]int{5, 4, 3, 2, 1}))
}
