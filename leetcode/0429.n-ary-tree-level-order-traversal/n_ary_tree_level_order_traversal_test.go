package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/lijinglin2019/algorithm-go/leetcode/common"
)

func TestLevelOrder(t *testing.T) {
	cases := []*Node{
		NodeExample1,
		{},
		nil,
	}
	results := [][][]int{
		{{0}, {1, 2, 3}, {4, 5, 6}},
		{{0}},
		nil,
	}

	for i := range cases {
		assert.Equal(t, results[i], levelOrder(cases[i]))
	}
}
