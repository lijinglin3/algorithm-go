package leetcode

//斜率，用最简约分子分母表示，为防止溢出，类型为int64
type slope struct {
	m int64 //分子
	n int64 //分母
}

func maxPoints(points [][]int) int {
	n := len(points)
	if n < 3 {
		return n
	}

	res := 0
	same := slope{0, 0}

	for i := 0; i < n-1; i++ {
		m := make(map[slope]int)
		temp := 0
		for j := i + 1; j < n; j++ {
			k := getSlope(points[i], points[j])
			m[k]++
			if k != same {
				temp = max(temp, m[k])
			}
		}
		res = max(res, temp+1+m[same])
	}
	return res
}

//根据两点求斜率
func getSlope(p1, p2 []int) slope {
	dx := int64(p1[0] - p2[0])
	dy := int64(p1[1] - p2[1])
	if dx == 0 && dy == 0 {
		return slope{int64(0), int64(0)}
	}
	if dx == 0 {
		return slope{int64(1), int64(0)}
	}
	if dy == 0 {
		return slope{int64(0), int64(1)}
	}
	if dx < 0 {
		dx = -dx
		dy = -dy
	}
	d := gcd(dx, dy)
	return slope{dx / d, dy / d}
}

// 最大公约数
func gcd(a, b int64) int64 {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
