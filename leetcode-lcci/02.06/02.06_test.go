package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.Equal(t, true, isPalindrome(nil))
	assert.Equal(t, true, isPalindrome(leetcode.NewListNode("[1]")))
	assert.Equal(t, true, isPalindrome(leetcode.NewListNode("[1,2,2,1]")))
	assert.Equal(t, false, isPalindrome(leetcode.NewListNode("[1,2]")))
}
