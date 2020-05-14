package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin2019/algorithm-go/leetcode"
)

func TestPreorder(t *testing.T) {
	cases := []*leetcode.Node{
		leetcode.NodeExample1,
		{},
		nil,
	}
	results := [][]int{
		{0, 1, 4, 5, 6, 2, 3},
		{0},
		nil,
	}

	for i := range cases {
		assert.Equal(t, results[i], Preorder(cases[i]))
		assert.Equal(t, results[i], PreorderByStack(cases[i]))
	}
}
