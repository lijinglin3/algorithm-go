package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShuffle(t *testing.T) {
	assert.Equal(t, []int{2, 3, 5, 4, 1, 7}, shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
}
