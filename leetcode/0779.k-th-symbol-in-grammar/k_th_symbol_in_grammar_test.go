package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthGrammar(t *testing.T) {
	assert.Equal(t, 0, kthGrammar(1, 1))
	assert.Equal(t, 0, kthGrammar(2, 1))
	assert.Equal(t, 1, kthGrammar(2, 2))
	assert.Equal(t, 1, kthGrammar(4, 5))
	assert.Equal(t, 0, kthGrammar(30, 434991989))
}
