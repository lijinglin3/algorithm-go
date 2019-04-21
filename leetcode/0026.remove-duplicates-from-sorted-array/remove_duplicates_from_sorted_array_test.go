package leetcode

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	var nums1 []int
	len1 := removeDuplicates(nums1)
	if !reflect.DeepEqual(nums1, []int{}) || len1 != 0 {
		t.Failed()
	}

	nums2 := []int{1}
	len2 := removeDuplicates(nums1)
	if !reflect.DeepEqual(nums2, []int{1}) || len2 != 1 {
		t.Failed()
	}

	nums3 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	len3 := removeDuplicates(nums3)
	if !reflect.DeepEqual(nums3, []int{0, 1, 2, 3, 4}) || len3 != 5 {
		t.Failed()
	}
}
