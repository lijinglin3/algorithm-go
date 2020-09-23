package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutation(t *testing.T) {
	assert.ElementsMatch(t, []string{""}, permutation(""))
	assert.ElementsMatch(t, []string{"abc", "acb", "bac", "bca", "cab", "cba"}, permutation("abc"))
}
