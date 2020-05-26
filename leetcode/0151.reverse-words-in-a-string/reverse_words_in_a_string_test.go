package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	assert.Equal(t, "", reverseWords("   "))
	assert.Equal(t, "hello", reverseWords("  hello  "))
	assert.Equal(t, "world! hello", reverseWords("  hello world!  "))
	assert.Equal(t, "blue is sky the", reverseWords("the sky is blue"))

	assert.Equal(t, "", reverseWords2("   "))
	assert.Equal(t, "hello", reverseWords2("  hello  "))
	assert.Equal(t, "world! hello", reverseWords2("  hello world!  "))
	assert.Equal(t, "blue is sky the", reverseWords2("the sky is blue"))
}
