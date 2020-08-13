package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToHex(t *testing.T) {
	assert.Equal(t, "1a", toHex(26))
	assert.Equal(t, "ffffffff", toHex(-1))
}
