package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	c := Constructor()
	str := "[1,2,3,null,null,4,5]"
	assert.Equal(t, str, c.serialize(c.deserialize(str)))
	assert.Equal(t, "[]", c.serialize(c.deserialize("[]")))
}
