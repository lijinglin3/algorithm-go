package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHammingWeight(t *testing.T) {
	assert.Equal(t, 3, hammingWeight(00000000000000000000000000001011))
}
