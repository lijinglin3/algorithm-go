package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin2019/algorithm-go/leetcode"
)

func TestHasCycle(t *testing.T) {
	listNode := leetcode.NewListNode("[1, 2, 3, 4]")
	assert.Equal(t, false, hasCycle(listNode))

	listNode = leetcode.NewListNode("[1, 2, 3, 4, 5]")
	assert.Equal(t, false, hasCycle(listNode))

	listNode.Next.Next.Next.Next.Next = listNode.Next.Next
	assert.Equal(t, true, hasCycle(listNode))
}
