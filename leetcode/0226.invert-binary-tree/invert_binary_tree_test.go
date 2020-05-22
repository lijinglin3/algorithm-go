package leetcode

import (
	"testing"

	"github.com/lijinglin2019/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestInvertTree(t *testing.T) {
	assert.Equal(t, leetcode.NewTreeNode("[4,7,2,9,6,3,1]"), invertTree(leetcode.NewTreeNode("[4,2,7,1,3,6,9]")))
}
