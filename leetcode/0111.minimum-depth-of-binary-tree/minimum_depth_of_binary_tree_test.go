package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/lijinglin2019/algorithm-go/leetcode/common"
)

func TestMinDepth(t *testing.T) {
	cases := []*TreeNode{
		TreeNodeExample1,
		{},
		nil,
	}
	results := []int{
		2,
		1,
		0,
	}

	for i := range cases {
		assert.Equal(t, results[i], MinDepth(cases[i]))
	}
}
