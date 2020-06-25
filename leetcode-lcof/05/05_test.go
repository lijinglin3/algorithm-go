package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceSpace(t *testing.T) {
	assert.Equal(t, "We%20are%20happy.", replaceSpace("We are happy."))
}
