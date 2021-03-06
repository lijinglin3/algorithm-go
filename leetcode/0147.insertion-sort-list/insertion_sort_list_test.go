package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin3/algorithm-go/leetcode"
)

func TestInsertionSortList(t *testing.T) {
	assert.Equal(t, (*leetcode.ListNode)(nil), insertionSortList(nil))
	assert.Equal(t, leetcode.NewListNode("[-1,0,3,4,5]"), insertionSortList(leetcode.NewListNode("[-1,5,3,4,0]")))
}
