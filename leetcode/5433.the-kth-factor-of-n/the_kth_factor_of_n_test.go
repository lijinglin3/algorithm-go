package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthFactor(t *testing.T) {
	assert.Equal(t, 3, kthFactor(12, 3))
	assert.Equal(t, -1, kthFactor(12, 9))
}
