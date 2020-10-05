package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestFindMode(t *testing.T) {
	assert.Equal(t, []int{2}, findMode(leetcode.NewTreeNode("[1,null,2,2]")))
}
