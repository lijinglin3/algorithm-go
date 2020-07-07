package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHammingWeight(t *testing.T) {
	assert.Equal(t, 3, hammingWeight(11))
	assert.Equal(t, 31, hammingWeight(4294967293))
}
