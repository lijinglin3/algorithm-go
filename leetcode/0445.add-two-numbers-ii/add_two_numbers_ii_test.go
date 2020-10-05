package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	example1 := "[9, 8, 7, 6, 5, 4, 3, 2, 1]"
	example2 := "[8, 7, 6, 5, 4, 3, 2, 1]"
	zero := "[0]"

	result := "[1, 0, 7, 5, 3, 0, 8, 6, 4, 2]"
	assert.Equal(t, leetcode.NewListNode(result),
		addTwoNumbers(leetcode.NewListNode(example1), leetcode.NewListNode(example2)),
	)

	assert.Equal(t, leetcode.NewListNode(result),
		addTwoNumbers(leetcode.NewListNode(example2), leetcode.NewListNode(example1)),
	)

	assert.Equal(t, leetcode.NewListNode(example1),
		addTwoNumbers(leetcode.NewListNode(example1), leetcode.NewListNode(zero)),
	)
	assert.Equal(t, leetcode.NewListNode(example1),
		addTwoNumbers(leetcode.NewListNode(zero), leetcode.NewListNode(example1)),
	)
}
