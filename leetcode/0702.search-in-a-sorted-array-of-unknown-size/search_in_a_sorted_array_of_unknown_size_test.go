package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	assert.Equal(t, 4, search(ArrayReader{[]int{-1, 0, 3, 5, 9, 12}}, 9))
	assert.Equal(t, -1, search(ArrayReader{[]int{-1, 0, 3, 5, 9, 12}}, 10))
}
