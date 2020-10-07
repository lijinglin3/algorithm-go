package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCuttingRope(t *testing.T) {
	assert.Equal(t, 1, cuttingRope(2))
	assert.Equal(t, 36, cuttingRope(10))
}
