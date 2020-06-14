package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	assert.Equal(t, []int{2}, intersection([]int{1, 2, 2, 1}, []int{2, 2}))
}
