package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin2019/algorithm-go/leetcode"
)

func TestRemoveNthFromEnd(t *testing.T) {
	listNode := "[1, 2, 3, 4, 5, 6]"
	assert.Equal(t, leetcode.NewListNode("[1, 2, 3, 4, 5]"),
		removeNthFromEnd(leetcode.NewListNode(listNode), 1))
	assert.Equal(t, leetcode.NewListNode("[1, 2, 3, 5, 6]"),
		removeNthFromEnd(leetcode.NewListNode(listNode), 3))
	assert.Equal(t, leetcode.NewListNode("[2, 3, 4, 5, 6]"),
		removeNthFromEnd(leetcode.NewListNode(listNode), 6))
}
