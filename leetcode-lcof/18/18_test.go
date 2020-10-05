package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestDeleteNode(t *testing.T) {
	assert.Equal(t, leetcode.NewListNode("[]"), deleteNode(nil, 1))
	assert.Equal(t, leetcode.NewListNode("[]"), deleteNode(leetcode.NewListNode("[1]"), 1))
	assert.Equal(t, leetcode.NewListNode("[4,5,9]"), deleteNode(leetcode.NewListNode("[4,5,1,9]"), 1))
	assert.Equal(t, leetcode.NewListNode("[4,5,1]"), deleteNode(leetcode.NewListNode("[4,5,1,9]"), 9))
}
