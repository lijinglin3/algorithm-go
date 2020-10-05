package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	example1 := "[1, 2, 3, 4, 5, 6, 7, 8, 9]"
	example2 := "[1, 2, 3, 4, 5, 6, 7, 8]"
	zero := "[0]"

	result := "[2, 4, 6, 8, 0, 3, 5, 7, 0, 1]"
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
