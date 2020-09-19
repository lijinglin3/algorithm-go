package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceWords(t *testing.T) {
	assert.Equal(t, "the cat was rat by the bat", replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))
	assert.Equal(t, "a a b c", replaceWords([]string{"a", "b", "c"}, "aadsfasf absbs bbab cadsfafs"))
	assert.Equal(t, "a a a a a a a a bbb baba a", replaceWords([]string{"a", "aa", "aaa", "aaaa"}, "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"))
}
