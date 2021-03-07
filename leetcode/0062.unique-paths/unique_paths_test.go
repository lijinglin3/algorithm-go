package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths(t *testing.T) {
	assert.Equal(t, 28, uniquePaths(3, 7))
}
