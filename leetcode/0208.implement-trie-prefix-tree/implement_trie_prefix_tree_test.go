package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	c := Constructor()

	c.Insert("apple")

	assert.Equal(t, true, c.Search("apple"))
	assert.Equal(t, false, c.Search("appleapple"))
	assert.Equal(t, false, c.Search("app"))

	assert.Equal(t, true, c.StartsWith("app"))
	assert.Equal(t, false, c.StartsWith("abc"))
	assert.Equal(t, true, c.StartsWith("apple"))
}
