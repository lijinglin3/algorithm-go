package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/lijinglin2019/algorithm-go/leetcode/common"
)

func TestIsSymmetric(t *testing.T) {
	cases := []*TreeNode{
		TreeNodeDecoder("[1, 2, 3, 4, 5]"),
		TreeNodeDecoder("[1, 2, 2, 3, null, 4]"),
		TreeNodeDecoder("[1, 2, 2, 3, 4, 4, 3]"),
		{},
		nil,
	}
	results := []bool{
		false,
		false,
		true,
		true,
		true,
	}

	for i := range cases {
		assert.Equal(t, results[i], IsSymmetric(cases[i]))
		assert.Equal(t, results[i], IsSymmetricByRecursive(cases[i]))
	}
}
