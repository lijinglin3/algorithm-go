package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRow(t *testing.T) {
	assert.Equal(t, []int{1}, getRow(0))
	assert.Equal(t, []int{1, 1}, getRow(1))
	assert.Equal(t, []int{1, 2, 1}, getRow(2))
}
