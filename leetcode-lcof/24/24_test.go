package leetcode

import (
	"testing"

	"github.com/lijinglin2019/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	assert.Equal(t, leetcode.NewListNode("[5,4,3,2,1]"), reverseList(leetcode.NewListNode("[1,2,3,4,5]")))
}
