package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstringKDistinct(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstringKDistinct("eceba", 2))
}
