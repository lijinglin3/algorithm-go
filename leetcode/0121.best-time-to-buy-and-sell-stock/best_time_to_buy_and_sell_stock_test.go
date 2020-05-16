package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	assert.Equal(t, 0, maxProfit([]int{7}))
	assert.Equal(t, 5, maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
