package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/lijinglin2019/algorithm-go/leetcode/common"
)

func TestMinDepth(t *testing.T) {
	cases := []*TreeNode{
		TreeNodeDecoder("[1, 2, 3, 4, 5]"),
		{},
		nil,
	}
	results := []int{
		2,
		1,
		0,
	}

	for i := range cases {
		assert.Equal(t, results[i], minDepth(cases[i]))
	}
}
