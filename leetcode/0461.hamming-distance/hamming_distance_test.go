package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHammingDistance(t *testing.T) {
	assert.Equal(t, 2, hammingDistance(1, 4))
	assert.Equal(t, 2, hammingDistance(4, 1))
}
