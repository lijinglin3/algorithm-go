package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPathCrossing(t *testing.T) {
	assert.Equal(t, false, isPathCrossing("NES"))
	assert.Equal(t, true, isPathCrossing("NESWW"))
}
