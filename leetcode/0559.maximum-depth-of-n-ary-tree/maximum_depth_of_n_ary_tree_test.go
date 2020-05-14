package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/lijinglin2019/algorithm-go/leetcode/common"
)

func TestMaxDepth(t *testing.T) {
	cases := []*Node{
		NodeExample1,
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
