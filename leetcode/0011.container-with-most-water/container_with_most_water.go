package leetcode

func maxArea(height []int) int {
	max := 0
	l, r := 0, len(height)-1
	for l < r {
		var area int
		if height[l] > height[r] {
			area = height[r] * (r - l)
			r--
		} else {
			area = height[l] * (r - l)
			l++
		}
		if area > max {
			max = area
		}
	}
	return max
}
