package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin3/algorithm-go/leetcode"
)

func TestPostorderTraversal(t *testing.T) {
	cases := []*leetcode.TreeNode{
		leetcode.NewTreeNode("[1, 2, 3, 4, 5]"),
		{},
		nil,
	}
	results := [][]int{
		{4, 5, 2, 3, 1},
		{0},
		nil,
	}

	for i := range cases {
		assert.Equal(t, results[i], postorderTraversal(cases[i]))
	}
}

func TestPostorderTraversalByStack(t *testing.T) {
	cases := []*leetcode.TreeNode{
		leetcode.NewTreeNode("[1, 2, 3, 4, 5]"),
		{},
		nil,
	}
	results := [][]int{
		{4, 5, 2, 3, 1},
		{0},
		nil,
	}

	for i := range cases {
		assert.Equal(t, results[i], postorderTraversalByStack(cases[i]))
	}
}
