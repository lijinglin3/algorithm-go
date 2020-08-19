package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSegments(t *testing.T) {
	assert.Equal(t, 5, countSegments("Hello, my name is John"))
}
