package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUnique(t *testing.T) {
	assert.Equal(t, true, isUnique("abcdefghi"))
	assert.Equal(t, false, isUnique("leetcode"))
}
