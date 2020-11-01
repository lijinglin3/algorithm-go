package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestGetIntersectionNode(t *testing.T) {
	example := leetcode.NewListNode("[1,8,4,5]")
	l1 := &leetcode.ListNode{Val: 4, Next: example}
	l2 := &leetcode.ListNode{Val: 5, Next: &leetcode.ListNode{Val: 0, Next: example}}

	assert.Equal(t, example, getIntersectionNode(l1, l2))
	assert.Equal(t, (*leetcode.ListNode)(nil), getIntersectionNode(l1, leetcode.NewListNode("[1,2,3,4,5]")))
	assert.Equal(t, (*leetcode.ListNode)(nil), getIntersectionNode(l1, nil))
	assert.Equal(t, (*leetcode.ListNode)(nil), getIntersectionNode(nil, nil))
}
