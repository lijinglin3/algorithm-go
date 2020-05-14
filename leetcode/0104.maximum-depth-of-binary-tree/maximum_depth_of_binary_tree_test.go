package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin2019/algorithm-go/leetcode"
)

func TestMaxDepth(t *testing.T) {
	cases := []*leetcode.TreeNode{
		leetcode.TreeNodeDecoder("[1, 2, 3, 4, 5]"),
		{},
		nil,
	}
	results := []int{
		3,
		1,
		0,
	}

	for i := range cases {
		assert.Equal(t, results[i], maxDepth(cases[i]))
		assert.Equal(t, results[i], maxDepthByBFS(cases[i]))
		assert.Equal(t, results[i], maxDepthByDFS(cases[i]))
	}
}
