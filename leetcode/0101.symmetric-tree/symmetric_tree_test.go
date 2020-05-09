package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/lijinglin2019/algorithm-go/leetcode/common"
)

func TestIsSymmetric(t *testing.T) {
	cases := []*TreeNode{
		TreeNodeExample1,
		TreeNodeExample2,
		TreeNodeExample3,
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
