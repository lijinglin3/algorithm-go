package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProduct(t *testing.T) {
	assert.Equal(t, 24, maxProduct([]int{1, 2, 3, 7, 5, 4}))
}
