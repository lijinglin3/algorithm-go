package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin2019/algorithm-go/leetcode"
)

func TestIsBalanced(t *testing.T) {
	cases := []*leetcode.TreeNode{
		leetcode.TreeNodeDecoder("[1, 2, 3, 4, 5]"),
		leetcode.TreeNodeDecoder("[1, 2, 2, 3, null, 4]"),
		leetcode.TreeNodeDecoder("[1, null, 2, null, 3]"),
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
		assert.Equal(t, results[i], isBalanced(cases[i]))
	}
}
