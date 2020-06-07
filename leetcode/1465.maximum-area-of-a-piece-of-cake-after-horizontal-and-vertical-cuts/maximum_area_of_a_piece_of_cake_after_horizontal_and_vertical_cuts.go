package leetcode

import "sort"

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	horizontalCuts = append(horizontalCuts, 0, h)
	verticalCuts = append(verticalCuts, 0, w)
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)

	maxH := 0
	for i := 1; i < len(horizontalCuts); i++ {
		delta := horizontalCuts[i] - horizontalCuts[i-1]
		if delta > maxH {
			maxH = delta
		}
	}

	maxW := 0
	for i := 1; i < len(verticalCuts); i++ {
		delta := verticalCuts[i] - verticalCuts[i-1]
		if delta > maxW {
			maxW = delta
		}
	}

	return (maxH * maxW) % (1000000000 + 7)
}
