package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidWordAbbreviation(t *testing.T) {
	assert.Equal(t, true, validWordAbbreviation("internationalization", "i12iz4n"))
	assert.Equal(t, true, validWordAbbreviation("internationalization", "i12iz5"))
	assert.Equal(t, false, validWordAbbreviation("internationalization", "i12iz4n0"))
	assert.Equal(t, false, validWordAbbreviation("internationalization", "i12iz9n"))
}
