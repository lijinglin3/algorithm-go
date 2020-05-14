package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin2019/algorithm-go/leetcode"
)

func TestSearchBST(t *testing.T) {
	bst := leetcode.TreeNodeDecoder("[4, 2, 7, 1, 3]")
	assert.Equal(t, leetcode.TreeNodeDecoder("[2, 1, 3]"), searchBST(bst, 2))
	assert.Equal(t, (*leetcode.TreeNode)(nil), searchBST(bst, 5))
	assert.Equal(t, (*leetcode.TreeNode)(nil), searchBST(nil, 0))
}
