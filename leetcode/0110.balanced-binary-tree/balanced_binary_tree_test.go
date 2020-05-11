package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/lijinglin2019/algorithm-go/leetcode/common"
)

func TestIsBalanced(t *testing.T) {
	cases := []*TreeNode{
		TreeNodeDecoder("[1, 2, 3, 4, 5]"),
		TreeNodeDecoder("[1, 2, 2, 3, null, 4]"),
		TreeNodeDecoder("[1, null, 2, null, 3]"),
		{},
		nil,
	}
	results := []bool{
		true,
		true,
		false,
		true,
		true,
	}

	for i := range cases {
		assert.Equal(t, results[i], IsBalanced(cases[i]))
	}
}
