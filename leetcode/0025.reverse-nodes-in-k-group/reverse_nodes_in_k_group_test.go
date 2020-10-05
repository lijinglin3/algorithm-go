package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin3/algorithm-go/leetcode"
)

func TestReverseKGroup(t *testing.T) {
	list := "[1, 2, 3, 4, 5, 6, 7, 8]"

	assert.Equal(t, leetcode.NewListNode(list),
		reverseKGroup(leetcode.NewListNode(list), 1))
	assert.Equal(t, leetcode.NewListNode("[2, 1, 4, 3, 6, 5, 8, 7]"),
		reverseKGroup(leetcode.NewListNode(list), 2))
	assert.Equal(t, leetcode.NewListNode("[3, 2, 1, 6, 5, 4, 7, 8]"),
		reverseKGroup(leetcode.NewListNode(list), 3))
}
