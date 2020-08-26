package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinRemoveToMakeValid(t *testing.T) {
	assert.Equal(t, "lee(t(c)o)de", minRemoveToMakeValid("lee(t(c)o)de)"))
	assert.Equal(t, "()()", minRemoveToMakeValid("())()((("))
}
