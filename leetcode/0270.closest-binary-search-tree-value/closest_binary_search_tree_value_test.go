package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestClosestValue(t *testing.T) {
	assert.Equal(t, 4, closestValue(leetcode.NewTreeNode("[4,2,5,1,3]"), 3.714286))
	assert.Equal(t, 3, closestValue(leetcode.NewTreeNode("[36,13,37,4,20,null,48,1,5,17,33,43,49,0,2,null,9,14,18,22,34,40,46,null,null,null,null,null,3,7,11,null,16,null,19,21,27,null,35,39,42,45,47,null,null,6,8,10,12,15,null,null,null,null,null,26,32,null,null,38,null,41,null,44,null,null,null,null,null,null,null,null,null,null,null,null,null,24,null,28,null,null,null,null,null,null,null,23,25,null,29,null,null,null,null,null,31,30]"), 3.142857))
}
