package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonChars(t *testing.T) {
	assert.Equal(t, []string{"e", "l", "l"}, commonChars([]string{"bella", "label", "roller"}))
	assert.Equal(t, []string{"c", "o"}, commonChars([]string{"cool", "lock", "cook"}))
}
