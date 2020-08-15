package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	assert.Equal(t, true, isSubsequence("abc", "ahbgdc"))
	assert.Equal(t, false, isSubsequence("axc", "ahbgdc"))
}
