package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	assert.Equal(t, "", reverseWords(""))
	assert.Equal(t, "s'teL ekat edoCteeL tsetnoc", reverseWords("Let's take LeetCode contest"))
}
