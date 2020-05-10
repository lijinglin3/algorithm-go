package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/lijinglin2019/algorithm-go/leetcode/common"
)

func TestIsBalanced(t *testing.T) {
	cases := []*TreeNode{
		TreeNodeExample1,
		TreeNodeExample2,
		TreeNodeExample4,
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
