package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin2019/algorithm-go/leetcode"
)

func TestDetectCycle(t *testing.T) {
	listNode := leetcode.ListNodeDecoder("[1, 2, 3, 4]")
	assert.Equal(t, (*leetcode.ListNode)(nil), detectCycle(listNode))

	listNode = leetcode.ListNodeDecoder("[1, 2, 3, 4, 5]")
	assert.Equal(t, (*leetcode.ListNode)(nil), detectCycle(listNode))

	listNode.Next.Next.Next.Next.Next = listNode.Next.Next
	assert.Equal(t, listNode.Next.Next, detectCycle(listNode))
}
