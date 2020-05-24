package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	c := Constructor()
	assert.Equal(t, "", c.serialize(nil))
	assert.Equal(t, (*Node)(nil), c.deserialize(""))

	assert.Equal(t, `{"Val":0,"Children":null}`, c.serialize(&Node{}))
	assert.Equal(t, &Node{}, c.deserialize(`{"Val":0,"Children":null}`))
}
