package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/lijinglin2019/algorithm-go/leetcode/common"
)

func TestHasPathSum(t *testing.T) {
	cases := []*TreeNode{
		TreeNodeDecoder("[1, 2, 3, 4, 5]"),
		TreeNodeDecoder("[1, 2, 3, 4, 5]"),
		{},
		nil,
	}
	sums := []int{
		8,
		100,
		0,
		0,
	}
	results := []bool{
		true,
		false,
		true,
		false,
	}

	for i := range cases {
		assert.Equal(t, results[i], hasPathSum(cases[i], sums[i]))
	}
}
