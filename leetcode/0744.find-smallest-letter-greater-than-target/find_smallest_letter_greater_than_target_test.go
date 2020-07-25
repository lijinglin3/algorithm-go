package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextGreatestLetter(t *testing.T) {
	assert.Equal(t, uint8('c'), nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a'))
	assert.Equal(t, uint8('f'), nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c'))
	assert.Equal(t, uint8('f'), nextGreatestLetter([]byte{'c', 'f', 'j'}, 'd'))
	assert.Equal(t, uint8('j'), nextGreatestLetter([]byte{'c', 'f', 'j'}, 'f'))
	assert.Equal(t, uint8('c'), nextGreatestLetter([]byte{'c', 'f', 'j'}, 'j'))
	assert.Equal(t, uint8('c'), nextGreatestLetter([]byte{'c', 'f', 'j'}, 'z'))
}
