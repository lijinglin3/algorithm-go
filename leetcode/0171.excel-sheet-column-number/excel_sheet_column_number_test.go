package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTitleToNumber(t *testing.T) {
	assert.Equal(t, 1, titleToNumber("A"))
	assert.Equal(t, 28, titleToNumber("AB"))
}
