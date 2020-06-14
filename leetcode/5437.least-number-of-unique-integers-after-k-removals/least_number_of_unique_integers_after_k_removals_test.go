package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLeastNumOfUniqueInts(t *testing.T) {
	assert.Equal(t, 1, findLeastNumOfUniqueInts([]int{5, 5, 4}, 1))
}
