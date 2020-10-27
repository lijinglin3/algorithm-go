package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	assert.Equal(t, leetcode.NewListNode("[0,1]"),
		addTwoNumbers(leetcode.NewListNode("[5]"), leetcode.NewListNode("[5]")))
	assert.Equal(t, leetcode.NewListNode("[2,1,9]"),
		addTwoNumbers(leetcode.NewListNode("[7,1,6]"), leetcode.NewListNode("[5,9,2]")))
}
