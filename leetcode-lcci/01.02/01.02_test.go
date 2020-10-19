package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPermutation(t *testing.T) {
	assert.Equal(t, true, CheckPermutation("abc", "cba"))
	assert.Equal(t, false, CheckPermutation("abcd", "cba"))
	assert.Equal(t, false, CheckPermutation("abcd", "cbaa"))
}
