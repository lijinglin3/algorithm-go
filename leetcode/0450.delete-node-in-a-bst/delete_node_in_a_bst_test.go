package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin2019/algorithm-go/leetcode"
)

/*
					10
			5			11
		4		6			12
	2				9
		3		8
*/
func TestDeleteNode(t *testing.T) {
	example := "[10, 5, 11, 4, 6, null, 12, 2, null, null, 9, null, null, null, 3, 8]"

	assert.Equal(t, leetcode.NewTreeNode("[10, 5, 11, 4, 6, null, null, 2, null, null, 9, null, 3, 8]"),
		deleteNode(leetcode.NewTreeNode(example), 12))

	assert.Equal(t, leetcode.NewTreeNode("[10, 5, 11, 3, 6, null, 12, 2, null, null, 9, null, null, null, null, 8]"),
		deleteNode(leetcode.NewTreeNode(example), 4))

	assert.Equal(t, leetcode.NewTreeNode("[10, 5, 11, 4, 8, null, 12, 2, null, null, 9, null, null, null, 3]"),
		deleteNode(leetcode.NewTreeNode(example), 6))

	assert.Equal(t, (*leetcode.TreeNode)(nil),
		deleteNode(leetcode.NewTreeNode("[]"), 0))
}
