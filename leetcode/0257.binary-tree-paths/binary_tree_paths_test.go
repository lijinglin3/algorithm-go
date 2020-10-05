package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestBinaryTreePaths(t *testing.T) {
	assert.ElementsMatch(t, ([]string)(nil), binaryTreePaths(leetcode.NewTreeNode("[]")))
	assert.ElementsMatch(t, []string{"1->2->5", "1->3"}, binaryTreePaths(leetcode.NewTreeNode("[1,2,3,null,5]")))
}
