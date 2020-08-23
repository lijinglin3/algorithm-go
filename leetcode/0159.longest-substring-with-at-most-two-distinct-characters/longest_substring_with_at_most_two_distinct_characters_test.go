package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstringTwoDistinct(t *testing.T) {
	assert.Equal(t, 5, lengthOfLongestSubstringTwoDistinct("ccaabbb"))
}
