package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceSpaces(t *testing.T) {
	assert.Equal(t, "Mr%20John%20Smith", replaceSpaces("Mr John Smith    ", 13))
	assert.Equal(t, "%20%20%20%20%20", replaceSpaces("               ", 5))
}
