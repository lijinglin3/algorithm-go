package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin3/algorithm-go/leetcode"
)

func TestMaxDepth(t *testing.T) {
	cases := []*leetcode.Node{
		leetcode.NodeExample1,
		{},
		nil,
	}
	results := []int{
		3,
		1,
		0,
	}

	for i := range cases {
		assert.Equal(t, results[i], maxDepth(cases[i]))
	}
}
