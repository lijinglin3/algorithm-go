package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountAndSay(t *testing.T) {
	assert.Equal(t, "1211", countAndSay(4))
}
