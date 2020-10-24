package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneEditAway(t *testing.T) {
	assert.Equal(t, true, oneEditAway("abc", "abc"))
	assert.Equal(t, true, oneEditAway("abc", "abd"))
	assert.Equal(t, true, oneEditAway("pale", "ple"))
	assert.Equal(t, true, oneEditAway("pale", "pales"))
	assert.Equal(t, false, oneEditAway("pales", "pal"))
}
