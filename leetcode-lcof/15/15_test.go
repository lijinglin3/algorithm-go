package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHammingWeight(t *testing.T) {
	assert.Equal(t, 3, hammingWeight(0b00000000000000000000000000001011))
	assert.Equal(t, 31, hammingWeight(0b11111111111111111111111111111101))
}
