package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrangeCoins(t *testing.T) {
	assert.Equal(t, 1, arrangeCoins(1))
	assert.Equal(t, 2, arrangeCoins(5))
	assert.Equal(t, 3, arrangeCoins(6))
}
