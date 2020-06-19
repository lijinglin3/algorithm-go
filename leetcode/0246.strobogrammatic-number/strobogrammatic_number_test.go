package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsStrobogrammatic(t *testing.T) {
	assert.Equal(t, true, isStrobogrammatic(""))
	assert.Equal(t, true, isStrobogrammatic("0"))
	assert.Equal(t, false, isStrobogrammatic("6"))
	assert.Equal(t, true, isStrobogrammatic("69"))
	assert.Equal(t, false, isStrobogrammatic("10"))
}
