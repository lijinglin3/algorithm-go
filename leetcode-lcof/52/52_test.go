package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestGetIntersectionNode(t *testing.T) {
	list1 := leetcode.NewListNode("[1, 2, 3]")
	list2 := leetcode.NewListNode("[1, 2, 3, 4]")
	assert.Equal(t, (*leetcode.ListNode)(nil), getIntersectionNode(list1, list2))

	common := leetcode.NewListNode("[7, 8, 9]")
	assert.Equal(t, common, getIntersectionNode(common, common))
	assert.Equal(t, (*leetcode.ListNode)(nil), getIntersectionNode(common, (*leetcode.ListNode)(nil)))

	list1.Last().Next, list2.Last().Next = common, common
	assert.Equal(t, common, getIntersectionNode(list1, common))
	assert.Equal(t, common, getIntersectionNode(list1, list2))
	assert.Equal(t, common, getIntersectionNode(list2, list1))
}
