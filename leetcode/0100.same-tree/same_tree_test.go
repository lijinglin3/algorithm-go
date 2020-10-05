package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestIsSameTree(t *testing.T) {
	tree1 := leetcode.NewTreeNode("[1,2,3]")
	tree2 := leetcode.NewTreeNode("[1,2,4]")
	assert.Equal(t, true, isSameTree(tree1, tree1))
	assert.Equal(t, false, isSameTree(tree1, tree2))
}
