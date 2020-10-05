package leetcode

import (
	"testing"

	"github.com/lijinglin2019/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestMergeTrees(t *testing.T) {
	assert.Equal(t,
		leetcode.NewTreeNode("[3,4,5,5,4,null,7]"),
		mergeTrees(
			leetcode.NewTreeNode("[1,3,2,5]"),
			leetcode.NewTreeNode("[2,1,3,null,4,null,7]"),
		),
	)
}
