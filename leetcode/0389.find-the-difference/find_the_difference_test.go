package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTheDifference(t *testing.T) {
	assert.Equal(t, byte('e'), findTheDifference("abcd", "edcba"))
}
