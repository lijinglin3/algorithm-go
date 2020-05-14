package leetcode

import (
	"testing"

	. "github.com/lijinglin2019/algorithm-go/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestPreorderTraversal(t *testing.T) {
	cases := []*TreeNode{
		TreeNodeDecoder("[1, 2, 3, 4, 5]"),
		{},
		nil,
	}
	results := [][]int{
		{1, 2, 4, 5, 3},
		{0},
		nil,
	}

	for i := range cases {
		assert.Equal(t, results[i], preorderTraversal(cases[i]))
	}
}

func TestPreorderTraversalByStack(t *testing.T) {
	cases := []*TreeNode{
		TreeNodeDecoder("[1, 2, 3, 4, 5]"),
		{},
		nil,
	}
	results := [][]int{
		{1, 2, 4, 5, 3},
		{0},
		nil,
	}

	for i := range cases {
		assert.Equal(t, results[i], preorderTraversalByStack(cases[i]))
	}
}
