package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStrongest(t *testing.T) {
	assert.Equal(t, []int{5, 1}, getStrongest([]int{1, 2, 3, 4, 5}, 2))
	assert.Equal(t, []int{5, 1, 4, 2, 3}, getStrongest([]int{1, 2, 3, 4, 5}, 10))
}
