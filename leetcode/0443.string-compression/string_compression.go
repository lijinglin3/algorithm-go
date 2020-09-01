package leetcode

import "strconv"

func compress(chars []byte) int {
	var left, right, size = 0, 0, 0
	for ; right <= len(chars); right++ {
		if right == len(chars) || chars[right] != chars[left] {
			chars[size] = chars[left]
			size++
			if right-left > 1 {
				l := strconv.Itoa(right - left)
				for i := range l {
					chars[size] = l[i]
					size++
				}
			}
			left = right
		}
	}
	return size
}
