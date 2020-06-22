package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFolderNames(t *testing.T) {
	assert.Equal(t, []string{"pes", "fifa", "gta", "pes(2019)"}, getFolderNames([]string{"pes", "fifa", "gta", "pes(2019)"}))
	assert.Equal(t, []string{"gta", "gta(1)", "gta(2)", "avalon"}, getFolderNames([]string{"gta", "gta(1)", "gta", "avalon"}))
	assert.Equal(t, []string{"kaido", "kaido(1)", "kaido(2)", "kaido(1)(1)"}, getFolderNames([]string{"kaido", "kaido(1)", "kaido", "kaido(1)"}))
}
