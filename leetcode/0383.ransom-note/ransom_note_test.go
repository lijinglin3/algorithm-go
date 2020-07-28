package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanConstruct(t *testing.T) {
	assert.Equal(t, false, canConstruct("a", "b"))
	assert.Equal(t, false, canConstruct("aa", "ab"))
	assert.Equal(t, false, canConstruct("aab", "aa"))
	assert.Equal(t, true, canConstruct("aa", "aab"))
}
