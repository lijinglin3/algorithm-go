package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestRemoveElements(t *testing.T) {
	assert.Equal(t, leetcode.NewListNode("[]"), removeElements(leetcode.NewListNode("[]"), 0))
	assert.Equal(t, leetcode.NewListNode("[]"), removeElements(leetcode.NewListNode("[1]"), 1))
	assert.Equal(t, leetcode.NewListNode("[1,2,3,4,5]"), removeElements(leetcode.NewListNode("[1,2,3,4,5,6]"), 6))
}
