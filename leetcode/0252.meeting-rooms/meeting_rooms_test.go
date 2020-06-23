package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanAttendMeetings(t *testing.T) {
	assert.Equal(t, false, canAttendMeetings([][]int{{0, 30}, {5, 10}, {15, 20}}))
	assert.Equal(t, true, canAttendMeetings([][]int{{7, 10}, {2, 4}}))
}
