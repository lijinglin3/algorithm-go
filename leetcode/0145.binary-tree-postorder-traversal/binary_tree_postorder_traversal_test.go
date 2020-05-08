package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/lijinglin2019/algorithm-go/leetcode/common"
)

func TestPostorderTraversal(t *testing.T) {
	cases := []*TreeNode{
		TreeNodeExample1,
		{},
		nil,
	}
	results := [][]int{
		{4, 5, 2, 3, 1},
		{0},
		nil,
	}

	for i := range cases {
		assert.Equal(t, results[i], PostorderTraversal(cases[i]))
	}
}

func TestPostorderTraversalByStack(t *testing.T) {
	cases := []*TreeNode{
		TreeNodeExample1,
		{},
		nil,
	}
	results := [][]int{
		{4, 5, 2, 3, 1},
		{0},
		nil,
	}

	for i := range cases {
		assert.Equal(t, results[i], PostorderTraversalByStack(cases[i]))
	}
}
