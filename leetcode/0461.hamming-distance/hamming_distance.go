package leetcode

func hammingDistance(x int, y int) int {
	rx, ry := make([]int, 0), make([]int, 0)
	for x != 0 {
		rx = append(rx, x%2)
		x = x / 2
	}
	for y != 0 {
		ry = append(ry, y%2)
		y = y / 2
	}

	lenx, leny := len(rx), len(ry)
	count := 0
	min := 0
	if lenx > leny {
		min = leny
		for i := leny; i < lenx; i++ {
			if rx[i] == 1 {
				count++
			}
		}
	} else {
		min = lenx
		for i := lenx; i < leny; i++ {
			if ry[i] == 1 {
				count++
			}
		}
	}

	for i := 0; i < min; i++ {
		if rx[i] != ry[i] {
			count++
		}
	}

	return count
}
