package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	r := reader{data: []byte("abcdefghijklmn")}
	s := solution(r.read4)
	buf := make([]byte, 10)
	assert.Equal(t, 10, s(buf, 10))
	assert.Equal(t, []byte("abcdefghij"), buf)

	assert.Equal(t, 2, s(buf, 10))
	assert.Equal(t, []byte("mncdefghij"), buf)
}

type reader struct {
	data []byte
	off  int
}

func (r *reader) read4(buf []byte) int {
	delta := 4
	if len(r.data)-r.off < 4 {
		delta = len(r.data) - r.off
	}
	r.off += delta
	copy(buf, r.data[r.off-delta:r.off])
	return delta
}
