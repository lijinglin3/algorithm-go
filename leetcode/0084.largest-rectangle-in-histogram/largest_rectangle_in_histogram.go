package leetcode

func largestRectangleArea(heights []int) int {
	heights = append([]int{-2}, heights...)
	heights = append(heights, -1)
	size := len(heights)
	s := make([]int, 1, size)

	res := 0
	i := 1
	for i < len(heights) {
		if heights[s[len(s)-1]] < heights[i] {
			s = append(s, i)
			i++
			continue
		}
		res = max(res, heights[s[len(s)-1]]*(i-s[len(s)-2]-1))
		s = s[:len(s)-1]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
