package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovingCount(t *testing.T) {
	assert.Equal(t, 3, movingCount(2, 3, 1))
	assert.Equal(t, 1, movingCount(3, 1, 0))
}
