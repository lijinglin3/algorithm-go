package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestSwapPairs(t *testing.T) {
	assert.Equal(t, leetcode.NewListNode("[2,1,4,3]"), swapPairs(leetcode.NewListNode("[1,2,3,4]")))
	assert.Equal(t, leetcode.NewListNode("[2,1,4,3,5]"), swapPairs(leetcode.NewListNode("[1,2,3,4,5]")))
}
