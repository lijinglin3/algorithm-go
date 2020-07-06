package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLastRemaining(t *testing.T) {
	assert.Equal(t, 3, lastRemaining(5, 3))
}
