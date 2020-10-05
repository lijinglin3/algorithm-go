package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin3/algorithm-go/leetcode"
)

func TestInorderTraversal(t *testing.T) {
	cases := []*leetcode.TreeNode{
		leetcode.NewTreeNode("[1, 2, 3, 4, 5]"),
		{},
		nil,
	}
	results := [][]int{
		{4, 2, 5, 1, 3},
		{0},
		nil,
	}

	for i := range cases {
		assert.Equal(t, results[i], inorderTraversal(cases[i]))
	}
}

func TestInorderTraversalByStack(t *testing.T) {
	cases := []*leetcode.TreeNode{
		leetcode.NewTreeNode("[1, 2, 3, 4, 5]"),
		{},
		nil,
	}
	results := [][]int{
		{4, 2, 5, 1, 3},
		{0},
		nil,
	}

	for i := range cases {
		assert.Equal(t, results[i], inorderTraversalByStack(cases[i]))
	}
}
