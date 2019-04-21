package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums, val := []int{3, 2, 2, 3}, 3
	len1 := removeElement(nums, val)
	fmt.Println(nums, len1)
	if len1 != 2 {
		t.Fail()
	}
}
