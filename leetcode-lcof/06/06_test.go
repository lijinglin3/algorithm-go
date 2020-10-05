package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestReversePrint(t *testing.T) {
	assert.Equal(t, []int{}, reversePrint(nil))
	assert.Equal(t, []int{}, reversePrint(leetcode.NewListNode("[]")))
	assert.Equal(t, []int{4, 5, 2, 3, 1}, reversePrint(leetcode.NewListNode("[1,3,2,5,4]")))
	assert.Equal(t, []int{6, 4, 5, 2, 3, 1}, reversePrint(leetcode.NewListNode("[1,3,2,5,4,6]")))
}
