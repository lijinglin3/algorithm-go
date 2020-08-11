package leetcode

import (
	"testing"

	"github.com/lijinglin2019/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestLowestCommonAncestor(t *testing.T) {
	assert.Equal(t, leetcode.NewTreeNode("[3,5,1,6,2,0,8,null,null,7,4]"),
		lowestCommonAncestor(
			leetcode.NewTreeNode("[3,5,1,6,2,0,8,null,null,7,4]"),
			&leetcode.TreeNode{Val: 6},
			&leetcode.TreeNode{Val: 8},
		),
	)
}
