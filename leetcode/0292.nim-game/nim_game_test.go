package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanWinNim(t *testing.T) {
	assert.Equal(t, false, canWinNim(4))
	assert.Equal(t, true, canWinNim(5))
	assert.Equal(t, true, canWinNim(6))
	assert.Equal(t, true, canWinNim(7))
	assert.Equal(t, false, canWinNim(8))
}
