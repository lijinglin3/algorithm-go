package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestLevelOrderBottom(t *testing.T) {
	cases := []*leetcode.TreeNode{
		leetcode.NewTreeNode("[1, 2, 3, 4, 5]"),
		{},
		nil,
	}
	results := [][][]int{
		{{4, 5}, {2, 3}, {1}},
		{{0}},
		nil,
	}

	for i := range cases {
		assert.Equal(t, results[i], levelOrderBottom(cases[i]))
	}
}
