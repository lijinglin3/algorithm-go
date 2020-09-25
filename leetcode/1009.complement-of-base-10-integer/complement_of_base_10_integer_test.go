package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitwiseComplement(t *testing.T) {
	assert.Equal(t, 1, bitwiseComplement(0))
	assert.Equal(t, 2, bitwiseComplement(5))
}
