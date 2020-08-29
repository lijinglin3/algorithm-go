package leetcode

import (
	"testing"

	"github.com/lijinglin2019/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestFlipEquiv(t *testing.T) {
	assert.Equal(t, true, flipEquiv(
		leetcode.NewTreeNode("[1,2,3,4,5,6,null,null,null,7,8]"),
		leetcode.NewTreeNode("[1,3,2,null,6,4,5,null,null,null,null,8,7]"),
	))
}
