package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin3/algorithm-go/leetcode"
)

func TestSearchBST(t *testing.T) {
	bst := leetcode.NewTreeNode("[4, 2, 7, 1, 3]")
	assert.Equal(t, leetcode.NewTreeNode("[2, 1, 3]"), searchBST(bst, 2))
	assert.Equal(t, (*leetcode.TreeNode)(nil), searchBST(bst, 5))
	assert.Equal(t, (*leetcode.TreeNode)(nil), searchBST(nil, 0))
}
