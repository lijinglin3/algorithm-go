package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/lijinglin2019/algorithm-go/leetcode/common"
)

func TestPreorder(t *testing.T) {
	cases := []*Node{
		NodeExample1,
		{},
		nil,
	}
	results := [][]int{
		{4, 5, 6, 1, 2, 3, 0},
		{0},
		nil,
	}

	for i := range cases {
		assert.Equal(t, results[i], Postorder(cases[i]))
		assert.Equal(t, results[i], PostorderByStack(cases[i]))
	}
}
