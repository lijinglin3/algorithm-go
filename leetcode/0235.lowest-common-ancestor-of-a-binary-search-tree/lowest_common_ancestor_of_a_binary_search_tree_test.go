package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin2019/algorithm-go/leetcode"
)

func TestLowestleetcodeAncestor(t *testing.T) {
	assert.Equal(t, 6, lowestCommonAncestor(
		leetcode.TreeNodeDecoder("[6,2,8,0,4,7,9,null,null,3,5]"),
		&leetcode.TreeNode{Val: 2},
		&leetcode.TreeNode{Val: 8},
	).Val)

	assert.Equal(t, 2, lowestCommonAncestor(
		leetcode.TreeNodeDecoder("[6,2,8,0,4,7,9,null,null,3,5]"),
		&leetcode.TreeNode{Val: 2},
		&leetcode.TreeNode{Val: 5},
	).Val)
	assert.Equal(t, 4, lowestCommonAncestor(
		leetcode.TreeNodeDecoder("[6,2,8,0,4,7,9,null,null,3,5]"),
		&leetcode.TreeNode{Val: 3},
		&leetcode.TreeNode{Val: 5},
	).Val)
	assert.Nil(t, lowestCommonAncestor(
		(*leetcode.TreeNode)(nil),
		(*leetcode.TreeNode)(nil),
		(*leetcode.TreeNode)(nil),
	))
}
