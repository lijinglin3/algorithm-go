package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	bh := Constructor("leetcode.com")
	bh.Visit("google.com")
	bh.Visit("facebook.com")
	bh.Visit("youtube.com")

	assert.Equal(t, "facebook.com", bh.Back(1))
	assert.Equal(t, "google.com", bh.Back(1))
	assert.Equal(t, "facebook.com", bh.Forward(1))

	bh.Visit("linkedin.com")

	assert.Equal(t, "linkedin.com", bh.Forward(2))
	assert.Equal(t, "google.com", bh.Back(2))
	assert.Equal(t, "leetcode.com", bh.Back(7))
}
