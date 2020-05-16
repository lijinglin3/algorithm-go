package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseBits(t *testing.T) {
	assert.Equal(t, uint32(964176192), reverseBits(43261596))
}
