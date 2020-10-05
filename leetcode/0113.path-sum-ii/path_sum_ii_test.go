package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestPathSum(t *testing.T) {
	tree := leetcode.NewTreeNode("[5,4,8,11,null,13,4,7,2,null,null,5,1]")
	result := [][]int{
		{5, 4, 11, 2},
		{5, 8, 4, 5},
	}
	assert.ElementsMatch(t, ([][]int)(nil), pathSum(leetcode.NewTreeNode("[]"), 22))
	assert.ElementsMatch(t, result, pathSum(tree, 22))
	assert.ElementsMatch(t, result, pathSum2(tree, 22))
}
