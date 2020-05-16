package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin2019/algorithm-go/leetcode"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	cases := []*leetcode.TreeNode{
		leetcode.NewTreeNode("[1, 2, 3, 4, 5]"),
		{},
		nil,
	}
	results := []int{3, 0, 0}

	for i := range cases {
		assert.Equal(t, results[i], diameterOfBinaryTree(cases[i]))
	}
}
