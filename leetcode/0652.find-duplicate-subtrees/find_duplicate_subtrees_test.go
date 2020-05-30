package leetcode

import (
	"testing"

	"github.com/lijinglin2019/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestFindDuplicateSubtrees(t *testing.T) {
	assert.ElementsMatch(t, []*leetcode.TreeNode{
		leetcode.NewTreeNode("[4]"),
		leetcode.NewTreeNode("[2,4]"),
	}, findDuplicateSubtrees(leetcode.NewTreeNode("[1,2,3,4,null,2,4,null,null,4]")))
}
