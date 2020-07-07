package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 10, add(5, 5))
	assert.Equal(t, 5, add(5, 0))
	assert.Equal(t, 0, add(5, -5))
}
