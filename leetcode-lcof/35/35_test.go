package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyRandomList(t *testing.T) {
	var (
		n0 = &Node{Val: 0, Next: nil, Random: nil}
		n1 = &Node{Val: 1, Next: n0, Random: nil}
		n2 = &Node{Val: 2, Next: n1, Random: nil}
		n3 = &Node{Val: 3, Next: n2, Random: nil}
		n4 = &Node{Val: 4, Next: n3, Random: nil}
		n5 = &Node{Val: 5, Next: n4, Random: nil}
	)
	n5.Random, n4.Random, n3.Random, n2.Random, n1.Random, n0.Random = n1, n2, n3, nil, nil, n5

	assert.Equal(t, n5, copyRandomList(n5))
	assert.Equal(t, (*Node)(nil), copyRandomList(nil))
}
