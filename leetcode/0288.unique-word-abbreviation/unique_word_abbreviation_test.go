package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	abbr := Constructor([]string{"it", "deer", "door", "cake", "card"})
	assert.Equal(t, true, abbr.IsUnique(""))
	assert.Equal(t, true, abbr.IsUnique("it"))
	assert.Equal(t, true, abbr.IsUnique("at"))
	assert.Equal(t, true, abbr.IsUnique("cart"))
	assert.Equal(t, true, abbr.IsUnique("cake"))
	assert.Equal(t, true, abbr.IsUnique("make"))
	assert.Equal(t, false, abbr.IsUnique("deer"))
	assert.Equal(t, false, abbr.IsUnique("dear"))
	assert.Equal(t, false, abbr.IsUnique("cane"))
}
