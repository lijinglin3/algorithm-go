package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacterReplacement(t *testing.T) {
	assert.Equal(t, 0, characterReplacement("", 1))
	assert.Equal(t, 4, characterReplacement("AABABBA", 1))
}
