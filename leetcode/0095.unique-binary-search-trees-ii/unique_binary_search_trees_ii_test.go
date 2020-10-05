package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestGenerateTrees(t *testing.T) {
	assert.Equal(t, ([]*leetcode.TreeNode)(nil), generateTrees(0))
	assert.ElementsMatch(t, []*leetcode.TreeNode{
		leetcode.NewTreeNode("[1,null,3,2]"),
		leetcode.NewTreeNode("[3,2,null,1]"),
		leetcode.NewTreeNode("[3,1,null,null,2]"),
		leetcode.NewTreeNode("[2,1,3]"),
		leetcode.NewTreeNode("[1,null,2,null,3]"),
	}, generateTrees(3))
}
