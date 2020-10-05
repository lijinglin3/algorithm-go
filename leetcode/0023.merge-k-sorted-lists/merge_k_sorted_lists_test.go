package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestMergeKLists(t *testing.T) {
	assert.Equal(t, leetcode.NewListNode("[1,1,2,3,4,4,5,6]"), mergeKLists([]*leetcode.ListNode{
		leetcode.NewListNode("[1,4,5]"),
		leetcode.NewListNode("[1,3,4]"),
		leetcode.NewListNode("[2,6]"),
		leetcode.NewListNode("[]"),
	}))
}
