package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupStrings(t *testing.T) {
	assert.ElementsMatch(t, [][]string{
		{"abc", "bcd", "xyz"},
		{"az", "ba"},
		{"acef"},
		{"a", "z"},
	}, groupStrings([]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}))
}
