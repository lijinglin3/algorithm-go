package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumTrees(t *testing.T) {
	assert.Equal(t, 5, numTrees(3))
	assert.Equal(t, 28853896093406129, numTrees(100))
}
