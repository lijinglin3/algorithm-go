package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFlipedString(t *testing.T) {
	assert.Equal(t, true, isFlipedString("waterbottle", "erbottlewat"))
	assert.Equal(t, false, isFlipedString("aa", "aba"))
}
