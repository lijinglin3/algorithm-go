package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/lijinglin2019/algorithm-go/leetcode/common"
)

func TestMaxDepth(t *testing.T) {
	cases := []*TreeNode{
		TreeNodeExample1,
		{},
		nil,
	}
	results := []int{
		3,
		1,
		0,
	}

	for i := range cases {
		assert.Equal(t, results[i], MaxDepth(cases[i]))
		assert.Equal(t, results[i], MaxDepthByBFS(cases[i]))
		assert.Equal(t, results[i], MaxDepthByDFS(cases[i]))
	}
}
