package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicateNodes(t *testing.T) {
	assert.Equal(t, (*leetcode.ListNode)(nil), removeDuplicateNodes(nil))
	assert.Equal(t, leetcode.NewListNode("[1, 2, 3]"), removeDuplicateNodes(leetcode.NewListNode("[1, 2, 3, 3, 2, 1]")))
	assert.Equal(t, leetcode.NewListNode("[1, 2]"), removeDuplicateNodes(leetcode.NewListNode("[1, 1, 1, 1, 2]")))
}
