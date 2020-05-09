package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/lijinglin2019/algorithm-go/leetcode/common"
)

func TestHasPathSum(t *testing.T) {
	cases := []*TreeNode{
		TreeNodeExample1,
		TreeNodeExample1,
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
		assert.Equal(t, results[i], HasPathSum(cases[i], sums[i]))
	}
}
