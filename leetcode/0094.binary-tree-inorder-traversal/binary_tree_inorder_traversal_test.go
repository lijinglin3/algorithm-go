package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/lijinglin2019/algorithm-go/leetcode/common"
)

func TestInorderTraversal(t *testing.T) {
	cases := []*TreeNode{
		TreeNodeDecoder("[1, 2, 3, 4, 5]"),
		{},
		nil,
	}
	results := [][]int{
		{4, 2, 5, 1, 3},
		{0},
		nil,
	}

	for i := range cases {
		assert.Equal(t, results[i], InorderTraversal(cases[i]))
	}
}

func TestInorderTraversalByStack(t *testing.T) {
	cases := []*TreeNode{
		TreeNodeDecoder("[1, 2, 3, 4, 5]"),
		{},
		nil,
	}
	results := [][]int{
		{4, 2, 5, 1, 3},
		{0},
		nil,
	}

	for i := range cases {
		assert.Equal(t, results[i], InorderTraversalByStack(cases[i]))
	}
}
