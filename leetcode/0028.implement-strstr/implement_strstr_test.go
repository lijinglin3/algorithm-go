package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	assert.Equal(t, 2, strStr("hello", "ll"))
}
