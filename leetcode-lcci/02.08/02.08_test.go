package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestDetectCycle(t *testing.T) {
	assert.Equal(t, (*leetcode.ListNode)(nil), detectCycle(nil))

	example := leetcode.NewListNode("[1,2,3,4,5]")
	assert.Equal(t, (*leetcode.ListNode)(nil), detectCycle(example))
	example.Next.Next.Next.Next.Next = example.Next.Next
	assert.Equal(t, example.Next.Next, detectCycle(example))
}
